package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectExampleRaw struct {
	Manager estype.Field[ManagerRaw] `json:"manager"`
}

type ManagerRaw struct {
	Age  estype.Field[int32]   `json:"age"`
	Name estype.Field[NameRaw] `json:"name"`
}

type NameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}
