package generate

import (
	"fmt"

	"github.com/ngicks/elastic-type/mapping"
)

func Nested(
	p mapping.NestedParams,
	globalOpt GlobalOption,
	opts MapOption,
	currentPointer []string,
) (highLevelTy, rawTy []GeneratedType, err error) {
	// Ignore dynamic inheritance.
	if p.Dynamic != nil && (*p.Dynamic == "true" || *p.Dynamic == true) {
		tyName := capitalize(globalOpt.TypeNameGenerator.Gen(currentPointer))
		return []GeneratedType{
				{
					TyName: tyName,
					TyDef:  fmt.Sprintf("type %s map[string]any", tyName),
				},
			}, []GeneratedType{
				{
					TyName: tyName + "Raw",
					TyDef:  fmt.Sprintf("type %s map[string]any", tyName+"Raw"),
				},
			}, nil
	}
	return object(*p.Properties, globalOpt, opts, currentPointer)
}
