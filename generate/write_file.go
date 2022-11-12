package generate

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ngicks/type-param-common/iterator"
	"github.com/ngicks/type-param-common/set"
)

var formatters = []string{
	"gofumpt",
	"goimports",
}

func systemHas(formatter string) bool {
	cmd := exec.Command("gofumpt", "--help")
	if cmd.Err != nil {
		return false
	}
	_, err := cmd.Output()
	return err == nil
}

func applyFormatter(srcPath string, formatter string) (stdout []byte, err error) {
	return exec.Command(formatter, "-w", srcPath).Output()
}

func formatter() (func(srcPath string) error, error) {
	for _, v := range formatters {
		if systemHas(v) {
			return func(srcPath string) error {
				_, err := applyFormatter(srcPath, v)
				return err
			}, nil
		}
	}

	return nil, fmt.Errorf("formatters not available: %+v", formatters)
}

func WriteFile(highLevelTyPath, rawTyePath string, highLevelTy, rawTy []GeneratedType, packageName string) error {
	var err error
	formatter, err := formatter()
	if err != nil {
		return err
	}

	highImports := extractImports(highLevelTy)
	highDef := extractDef(highLevelTy)

	rawImports := extractImports(rawTy)
	rawDef := extractDef(rawTy)

	var file *os.File

	writeFile := func(outPath, imports, def string) error {
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

	err = writeFile(highLevelTyPath, highImports, highDef)
	if err != nil {
		return err
	}
	err = writeFile(rawTyePath, rawImports, rawDef)
	if err != nil {
		return err
	}

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
