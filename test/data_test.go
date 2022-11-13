package test_test

import (
	_ "embed"
	"encoding/json"
)

var (
	// TODO: all.json must be expanded to fill all values.
	//go:embed example/all.json
	allJSONBin []byte
	//go:embed example/object_dynamic_inheritance.json
	objectInheritanceJSONBin []byte
	//go:embed example/object_w_overlap.json
	objectWOverlapJSONBin []byte
	//go:embed example/object.json
	objectJSONBin []byte
)

var (
	allMappings               []byte
	objectInheritanceMappings []byte
	objectWOverlapMappings    []byte
	objectMappings            []byte
	testSettings              = map[string]any{
		"number_of_replicas": 0, // This prevents es from being yellow after creation of index. Only needed if es is single-node.
	}
)

func init() {
	type Tuple struct {
		A []byte
		B *[]byte
	}
	bins := []Tuple{
		{allJSONBin, &allMappings},
		{A: objectInheritanceJSONBin, B: &objectInheritanceMappings},
		{A: objectWOverlapJSONBin, B: &objectWOverlapMappings},
		{A: objectJSONBin, B: &objectMappings},
	}
	for _, tuple := range bins {
		var mm map[string]map[string]any

		var err error
		err = json.Unmarshal(tuple.A, &mm)
		if err != nil {
			panic(err)
		}

		indexSettings := map[string]any{
			"settings": testSettings,
			"mappings": getOne(mm)["mappings"],
		}

		*tuple.B, err = json.MarshalIndent(indexSettings, "", "     ")
		if err != nil {
			panic(err)
		}
	}
}
