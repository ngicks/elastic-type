package example

type ObjectWOverlap struct {
	Manager     *[]ObjectWOverlapManager     `json:"manager"`
	Subordinate *[]ObjectWOverlapSubordinate `json:"subordinate"`
}

type ObjectWOverlapManager struct {
	Age  *[]int32              `json:"age"`
	Name *[]ObjectWOverlapName `json:"name"`
}

type ObjectWOverlapName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}

type ObjectWOverlapSubordinate struct {
	Age  *[]int32                         `json:"age"`
	Name *[]ObjectWOverlapSubordinateName `json:"name"`
}

type ObjectWOverlapSubordinateName struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}
