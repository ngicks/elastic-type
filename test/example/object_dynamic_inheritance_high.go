package example

type ObjectDynamicInheritance struct {
	Manager *[]ObjectDynamicInheritanceManager `json:"manager"`
	Player  *[]ObjectDynamicInheritancePlayer  `json:"player"`
}

type ObjectDynamicInheritanceManager struct {
	Age  *[]int32                        `json:"age"`
	Name *[]ObjectDynamicInheritanceName `json:"name"`
}

type ObjectDynamicInheritanceName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

type ObjectDynamicInheritancePlayer map[string]any
