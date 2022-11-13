package example

import (
	estype "github.com/ngicks/elastic-type/es_type"
)

type ObjectWOverlap struct {
	Manager     *[]ObjectWOverlapManager     `json:"manager"`
	Subordinate *[]ObjectWOverlapSubordinate `json:"subordinate"`
}

func (t ObjectWOverlap) ToRaw() ObjectWOverlapRaw {
	return ObjectWOverlapRaw{
		Manager: estype.MapField(estype.NewField(t.Manager, false), func(v ObjectWOverlapManager) ObjectWOverlapManagerRaw {
			return v.ToRaw()
		}),
		Subordinate: estype.MapField(estype.NewField(t.Subordinate, false), func(v ObjectWOverlapSubordinate) ObjectWOverlapSubordinateRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectWOverlapManager struct {
	Age  *[]int32              `json:"age"`
	Name *[]ObjectWOverlapName `json:"name"`
}

func (t ObjectWOverlapManager) ToRaw() ObjectWOverlapManagerRaw {
	return ObjectWOverlapManagerRaw{
		Age: estype.NewField(t.Age, false),
		Name: estype.MapField(estype.NewField(t.Name, false), func(v ObjectWOverlapName) ObjectWOverlapNameRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectWOverlapName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

func (t ObjectWOverlapName) ToRaw() ObjectWOverlapNameRaw {
	return ObjectWOverlapNameRaw{
		First: estype.NewField(t.First, false),
		Last:  estype.NewField(t.Last, false),
	}
}

type ObjectWOverlapSubordinate struct {
	Age  *[]int32                         `json:"age"`
	Name *[]ObjectWOverlapSubordinateName `json:"name"`
}

func (t ObjectWOverlapSubordinate) ToRaw() ObjectWOverlapSubordinateRaw {
	return ObjectWOverlapSubordinateRaw{
		Age: estype.NewField(t.Age, false),
		Name: estype.MapField(estype.NewField(t.Name, false), func(v ObjectWOverlapSubordinateName) ObjectWOverlapSubordinateNameRaw {
			return v.ToRaw()
		}),
	}
}

type ObjectWOverlapSubordinateName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

func (t ObjectWOverlapSubordinateName) ToRaw() ObjectWOverlapSubordinateNameRaw {
	return ObjectWOverlapSubordinateNameRaw{
		First: estype.NewField(t.First, false),
		Last:  estype.NewField(t.Last, false),
	}
}
