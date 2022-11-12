package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectExampleRaw struct {
	Manager estype.Field[ObjectExampleManagerRaw] `json:"manager"`
}

type ObjectExampleManagerRaw struct {
	Age  estype.Field[int32]                `json:"age" esjson:"single"`
	Name estype.Field[ObjectExampleNameRaw] `json:"name" esjson:"single"`
}

type ObjectExampleNameRaw struct {
	First estype.Field[string] `json:"first" esjson:"single"`
	Last  estype.Field[string] `json:"last"`
}
