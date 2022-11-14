package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectExample struct {
	Manager *[]ObjectExampleManager `json:"manager"`
}

func (t ObjectExample) ToRaw() ObjectExampleRaw {
	return ObjectExampleRaw{
		Manager: estype.MapField(estype.NewField(t.Manager), func(v ObjectExampleManager) ObjectExampleManagerRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectExampleManager struct {
	Age  int32             `json:"age"`
	Name ObjectExampleName `json:"name"`
}

func (t ObjectExampleManager) ToRaw() ObjectExampleManagerRaw {
	return ObjectExampleManagerRaw{
		Age: estype.NewFieldSingleValue(t.Age),
		Name: estype.MapField(estype.NewFieldSingleValue(t.Name), func(v ObjectExampleName) ObjectExampleNameRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectExampleName struct {
	First string   `json:"first"`
	Last  []string `json:"last"`
}

func (t ObjectExampleName) ToRaw() ObjectExampleNameRaw {
	return ObjectExampleNameRaw{
		First: estype.NewFieldSingleValue(t.First),
		Last:  estype.NewFieldSlice(t.Last, false),
	}
}
