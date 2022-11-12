package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectWOverlapRaw struct {
	Manager     estype.Field[ObjectWOverlapManagerRaw]     `json:"manager"`
	Subordinate estype.Field[ObjectWOverlapSubordinateRaw] `json:"subordinate"`
}

type ObjectWOverlapManagerRaw struct {
	Age  estype.Field[int32]                 `json:"age"`
	Name estype.Field[ObjectWOverlapNameRaw] `json:"name"`
}

type ObjectWOverlapNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

type ObjectWOverlapSubordinateRaw struct {
	Age  estype.Field[int32]                            `json:"age"`
	Name estype.Field[ObjectWOverlapSubordinateNameRaw] `json:"name"`
}

type ObjectWOverlapSubordinateNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}
