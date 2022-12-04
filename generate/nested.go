package generate

import (
	"fmt"

	"github.com/ngicks/elastic-type/mapping"
)

func Nested(
	p mapping.NestedParams,
	globalOpt GlobalOption,
	opts MapOption,
	fieldNames []string,
	dynamicContext mapping.Dynamic,
) (highLevelTy, rawTy, testDef []GeneratedType, err error) {
	newDynamic := mapping.OverlayDynamic(dynamicContext, p.Dynamic)

	// Ignore dynamic == "runtime" or dynamic == "false"
	if mapping.DynamicIsTrue(newDynamic) {
		tyName := capitalize(globalOpt.TypeNameGenerator.Gen(fieldNames))
		return []GeneratedType{
				{
					TyName: tyName,
					TyDef:  fmt.Sprintf("type %s map[string]any", tyName),
				},
			},
			[]GeneratedType{
				{
					TyName: tyName + "Raw",
					TyDef:  fmt.Sprintf("type %s map[string]any", tyName+"Raw"),
				},
			},
			[]GeneratedType{},
			nil
	}
	return object(*p.Properties, globalOpt, opts, fieldNames, newDynamic)
}
