package test

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
	AllMappings               []byte
	ObjectInheritanceMappings []byte
	ObjectWOverlapMappings    []byte
	ObjectMappings            []byte
	TestSettings              = map[string]any{
		"number_of_replicas": 0, // This prevents es from being yellow after creation of index. Only needed if es is single-node.
	}
)

func init() {
	type Tuple struct {
		A []byte
		B *[]byte
	}
	bins := []Tuple{
		{allJSONBin, &AllMappings},
		{A: objectInheritanceJSONBin, B: &ObjectInheritanceMappings},
		{A: objectWOverlapJSONBin, B: &ObjectWOverlapMappings},
		{A: objectJSONBin, B: &ObjectMappings},
	}
	for _, tuple := range bins {
		var mm map[string]map[string]any

		var err error
		err = json.Unmarshal(tuple.A, &mm)
		if err != nil {
			panic(err)
		}

		indexSettings := map[string]any{
			"settings": TestSettings,
			"mappings": getOne(mm)["mappings"],
		}

		*tuple.B, err = json.MarshalIndent(indexSettings, "", "     ")
		if err != nil {
			panic(err)
		}
	}
}

func getOne(v map[string]map[string]any) map[string]any {
	for _, v := range v {
		return v
	}
	return nil
}
