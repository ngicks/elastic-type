package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectDynamicInheritanceRaw struct {
	Manager estype.Field[ObjectDynamicInheritanceManagerRaw] `json:"manager"`
	Player  estype.Field[ObjectDynamicInheritancePlayerRaw]  `json:"player"`
}

func (r ObjectDynamicInheritanceRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectDynamicInheritanceRaw) ToPlain() ObjectDynamicInheritance {
	return ObjectDynamicInheritance{
		Manager: estype.MapField(t.Manager, func(v ObjectDynamicInheritanceManagerRaw) ObjectDynamicInheritanceManager {
			return v.ToPlain()
		}).Value(),
		Player: estype.MapField(t.Player, func(v ObjectDynamicInheritancePlayerRaw) ObjectDynamicInheritancePlayer {
			return v.ToPlain()
		}).Value(),
	}
}

type ObjectDynamicInheritanceManagerRaw struct {
	Age  estype.Field[int32]                           `json:"age"`
	Name estype.Field[ObjectDynamicInheritanceNameRaw] `json:"name"`
}

func (r ObjectDynamicInheritanceManagerRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectDynamicInheritanceManagerRaw) ToPlain() ObjectDynamicInheritanceManager {
	return ObjectDynamicInheritanceManager{
		Age: t.Age.Value(),
		Name: estype.MapField(t.Name, func(v ObjectDynamicInheritanceNameRaw) ObjectDynamicInheritanceName {
			return v.ToPlain()
		}).Value(),
	}
}

type ObjectDynamicInheritanceNameRaw struct {
	First estype.Field[string] `json:"first"`
	Last  estype.Field[string] `json:"last"`
}

func (r ObjectDynamicInheritanceNameRaw) MarshalJSON() ([]byte, error) {
	return estype.MarshalFieldsJSON(r)
}

func (t ObjectDynamicInheritanceNameRaw) ToPlain() ObjectDynamicInheritanceName {
	return ObjectDynamicInheritanceName{
		First: t.First.Value(),
		Last:  t.Last.Value(),
	}
}

type ObjectDynamicInheritancePlayerRaw map[string]estype.Field[any]

func (t ObjectDynamicInheritancePlayerRaw) ToPlain() ObjectDynamicInheritancePlayer {
	out := ObjectDynamicInheritancePlayer{}
	for k, v := range t {
		out[k] = v.ValueZero()
	}
	return out
}
