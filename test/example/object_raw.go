package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectExampleRaw struct {
	Manager estype.Field[ObjectExampleManagerRaw] `json:"manager"`
}

func (r ObjectExampleRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectExampleRaw) ToPlain() ObjectExample {
	return ObjectExample{
		Manager: estype.MapField(t.Manager, func(v ObjectExampleManagerRaw) ObjectExampleManager {
			return v.ToPlain()
		}).Value(),
	}
}

type ObjectExampleManagerRaw struct {
	Age  estype.Field[int32]                `json:"age" esjson:"single"`
	Name estype.Field[ObjectExampleNameRaw] `json:"name" esjson:"single"`
}

func (r ObjectExampleManagerRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectExampleManagerRaw) ToPlain() ObjectExampleManager {
	return ObjectExampleManager{
		Age: t.Age.ValueSingleZero(),
		Name: estype.MapField(t.Name, func(v ObjectExampleNameRaw) ObjectExampleName {
			return v.ToPlain()
		}).ValueSingleZero(),
	}
}

type ObjectExampleNameRaw struct {
	First estype.Field[string] `json:"first" esjson:"single"`
	Last  estype.Field[string] `json:"last"`
}

func (r ObjectExampleNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectExampleNameRaw) ToPlain() ObjectExampleName {
	return ObjectExampleName{
		First: t.First.ValueSingleZero(),
		Last:  t.Last.ValueZero(),
	}
}
