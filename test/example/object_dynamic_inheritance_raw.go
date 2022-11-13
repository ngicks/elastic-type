package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectDynamicInheritanceRaw struct {
	Manager estype.Field[ObjectDynamicInheritanceManagerRaw] `json:"manager"`
	Player  estype.Field[ObjectDynamicInheritancePlayerRaw]  `json:"player"`
}

type ObjectDynamicInheritanceManagerRaw struct {
	Age  estype.Field[int32]                           `json:"age"`
	Name estype.Field[ObjectDynamicInheritanceNameRaw] `json:"name"`
}

type ObjectDynamicInheritanceNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

type ObjectDynamicInheritancePlayerRaw map[string]any
