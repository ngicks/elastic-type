package generate

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ngicks/type-param-common/iterator"
	"github.com/ngicks/type-param-common/set"
)

var possibleFormatters = [][][]string{
	// It is really slower when using gopls. The combo of goimports and gofumpt is preferred here.
	{{"goimports", "-w"}, {"gofumpt", "-w"}},
	{{"gopls", "imports", "-w"}, {"gopls", "format", "-w"}},
}

type applyFormat = func(srcPath string) error

var formatCommands []applyFormat

func initializeFormatterCommands() {
	if formatCommands != nil {
		return
	}

	for _, commands := range possibleFormatters {
		systemHasAll := false
		for _, command := range commands {
			_, err := exec.LookPath(command[0])
			systemHasAll = err == nil
		}
		if systemHasAll {
			formatCommands = append(formatCommands, buildFormatCommand(commands))
		}
	}
}

func buildFormatCommand(commands [][]string) applyFormat {
	return func(srcPath string) error {
		for _, command := range commands {
			name := command[0]
			args := command[1:]
			err := exec.Command(name, append(args, srcPath)...).Run()
			if err != nil {
				return err
			}
		}
		return nil
	}
}

var ErrNoFormatterAvailable = errors.New(
	"no formatter available. " +
		"There must be at least an available formatters combination",
)

func WriteFile(highLevelTyPath, rawTyePath string, highLevelTy, rawTy []GeneratedType, packageName string) error {
	initializeFormatterCommands()

	if len(formatCommands) == 0 {
		return fmt.Errorf("%w: %v", ErrNoFormatterAvailable, possibleFormatters)
	}

	highImports := extractImports(highLevelTy)
	highDef := extractDef(highLevelTy)

	rawImports := extractImports(rawTy)
	rawDef := extractDef(rawTy)

	writeFile := func(outPath, imports, def string) error {
		var file *os.File
		var err error
		file, err = os.Create(outPath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, strings.NewReader(fmt.Sprintf("package %s\n\n", packageName)))
		if err != nil {
			return err
		}
		_, err = io.Copy(file, strings.NewReader(fmt.Sprintf("import (\n%s\n)", imports)))
		if err != nil {
			return err
		}
		_, err = io.Copy(file, strings.NewReader(def))
		if err != nil {
			return err
		}
		return nil
	}

	var err error
	err = writeFile(highLevelTyPath, highImports, highDef)
	if err != nil {
		return err
	}
	err = writeFile(rawTyePath, rawImports, rawDef)
	if err != nil {
		return err
	}

	formatter := formatCommands[0]
	err = formatter(highLevelTyPath)
	if err != nil {
		return err
	}
	err = formatter(rawTyePath)
	if err != nil {
		return err
	}

	return nil
}

func extractImports(ty []GeneratedType) string {
	imports := iterator.Map[GeneratedType](
		iterator.FromSlice(ty),
		func(gt GeneratedType) []string {
			return gt.Imports
		},
	)

	importsSet := iterator.Fold[[]string](
		imports,
		func(accumulator *set.Set[string], next []string) *set.Set[string] {
			iterator.FromSlice(next).
				Map(strings.TrimSpace).
				Exclude(func(s string) bool { return s == "" }).
				ForEach(func(s string) {
					accumulator.Add(s)
				})
			return accumulator
		},
		set.New[string](),
	)

	return strings.Join(importsSet.Values().Collect(), "\n")
}

func extractDef(ty []GeneratedType) string {
	defs := iterator.Map[GeneratedType](
		iterator.FromSlice(ty),
		func(gt GeneratedType) string {
			return gt.TyDef
		},
	).Exclude(func(s string) bool { return s == "" }).Collect()

	return strings.Join(defs, "\n")
}
