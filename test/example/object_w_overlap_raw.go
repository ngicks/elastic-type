package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectWOverlapRaw struct {
	Manager     estype.Field[ObjectWOverlapManagerRaw]     `json:"manager"`
	Subordinate estype.Field[ObjectWOverlapSubordinateRaw] `json:"subordinate"`
}

func (r ObjectWOverlapRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectWOverlapRaw) ToPlain() ObjectWOverlap {
	return ObjectWOverlap{
		Manager: estype.MapField(t.Manager, func(v ObjectWOverlapManagerRaw) ObjectWOverlapManager {
			return v.ToPlain()
		}).Value(),
		Subordinate: estype.MapField(t.Subordinate, func(v ObjectWOverlapSubordinateRaw) ObjectWOverlapSubordinate {
			return v.ToPlain()
		}).Value(),
	}
}

type ObjectWOverlapManagerRaw struct {
	Age  estype.Field[int32]                 `json:"age"`
	Name estype.Field[ObjectWOverlapNameRaw] `json:"name"`
}

func (r ObjectWOverlapManagerRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectWOverlapManagerRaw) ToPlain() ObjectWOverlapManager {
	return ObjectWOverlapManager{
		Age: t.Age.Value(),
		Name: estype.MapField(t.Name, func(v ObjectWOverlapNameRaw) ObjectWOverlapName {
			return v.ToPlain()
		}).Value(),
	}
}

type ObjectWOverlapNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

func (r ObjectWOverlapNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectWOverlapNameRaw) ToPlain() ObjectWOverlapName {
	return ObjectWOverlapName{
		First: t.First.Value(),
		Last:  t.Last.Value(),
	}
}

type ObjectWOverlapSubordinateRaw struct {
	Age  estype.Field[int32]                            `json:"age"`
	Name estype.Field[ObjectWOverlapSubordinateNameRaw] `json:"name"`
}

func (r ObjectWOverlapSubordinateRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectWOverlapSubordinateRaw) ToPlain() ObjectWOverlapSubordinate {
	return ObjectWOverlapSubordinate{
		Age: t.Age.Value(),
		Name: estype.MapField(t.Name, func(v ObjectWOverlapSubordinateNameRaw) ObjectWOverlapSubordinateName {
			return v.ToPlain()
		}).Value(),
	}
}

type ObjectWOverlapSubordinateNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

func (r ObjectWOverlapSubordinateNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectWOverlapSubordinateNameRaw) ToPlain() ObjectWOverlapSubordinateName {
	return ObjectWOverlapSubordinateName{
		First: t.First.Value(),
		Last:  t.Last.Value(),
	}
}
