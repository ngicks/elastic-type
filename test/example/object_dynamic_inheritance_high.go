package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectDynamicInheritance struct {
	Manager *[]ObjectDynamicInheritanceManager `json:"manager"`
	Player  *[]ObjectDynamicInheritancePlayer  `json:"player"`
}

func (t ObjectDynamicInheritance) ToRaw() ObjectDynamicInheritanceRaw {
	return ObjectDynamicInheritanceRaw{
		Manager: estype.MapField(estype.NewField(t.Manager), func(v ObjectDynamicInheritanceManager) ObjectDynamicInheritanceManagerRaw {
			return v.ToRaw()
		}),
		Player: estype.MapField(estype.NewField(t.Player), func(v ObjectDynamicInheritancePlayer) ObjectDynamicInheritancePlayerRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectDynamicInheritanceManager struct {
	Age  *[]int32                        `json:"age"`
	Name *[]ObjectDynamicInheritanceName `json:"name"`
}

func (t ObjectDynamicInheritanceManager) ToRaw() ObjectDynamicInheritanceManagerRaw {
	return ObjectDynamicInheritanceManagerRaw{
		Age: estype.NewField(t.Age),
		Name: estype.MapField(estype.NewField(t.Name), func(v ObjectDynamicInheritanceName) ObjectDynamicInheritanceNameRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectDynamicInheritanceName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

func (t ObjectDynamicInheritanceName) ToRaw() ObjectDynamicInheritanceNameRaw {
	return ObjectDynamicInheritanceNameRaw{
		First: estype.NewField(t.First),
		Last:  estype.NewField(t.Last),
	}
}

type ObjectDynamicInheritancePlayer map[string][]any

func (t ObjectDynamicInheritancePlayer) ToRaw() ObjectDynamicInheritancePlayerRaw {
	out := ObjectDynamicInheritancePlayerRaw{}
	for k, v := range t {
		out[k] = estype.NewFieldSlice(v, false)
	}
	return out
}
