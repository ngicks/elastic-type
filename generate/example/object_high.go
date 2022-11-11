package example

type ObjectExample struct {
	Manager *[]Manager `json:"manager"`
}

type Manager struct {
	Age  *[]int32 `json:"age"`
	Name *[]Name  `json:"name"`
}

type Name struct {
	First *[]string `json:"first"`
	Last  *[]string `json:"last"`
}
