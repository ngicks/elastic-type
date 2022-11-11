package generate

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ngicks/elastic-type/mapping"
	"github.com/ngicks/type-param-common/iterator"
	"github.com/ngicks/type-param-common/set"
)

type GeneratedType struct {
	TyName  string
	TyDef   string
	Imports []string
	Option  FieldOption
}

// Generate generates Go struct types from an Elasticsearch mapping.
//
// It generates 2 type implementations, high level one and raw one.
// The high level type is like a plain go struct, where a field that should be:
//   - many, is []T
//   - single, is T
//   - optional, is *T (or *[]T)
//   - required, is T (or []T)
//
// The raw type is, on the other hand,
// a type that can be directly unmarshalled from an Elasticsearch-stored json.
// All fields in the type can be null, undefined, T or []T.
// This is done with help of estype.Field[T any].
//
// The Elasticsearch (or Apache Lucene) is so _elastic_ that you can store every above variants of T.
// opts here is meta data which will be used to build an assumption about the data format you are to store,
// like optional|required or single|many, by our own.
// Keys of opts must be mapping property names. ChildOption will be only used for Object or Nested, for other types simply ignored.
//
// Always len(highLevenTy) == len(rawTy).
func Generate(props mapping.Properties, tyName string, globalOpt GlobalOption, opts MapOption) (highLevelTy, rawTy []GeneratedType, err error) {
	if opts == nil {
		opts = MapOption{}
	}
	return object(props, globalOpt, opts, []string{tyName})
}

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

		_, err = io.Copy(file, strings.NewReader(fmt.Sprintf("package %s", packageName)))
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
