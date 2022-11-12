package example

type ObjectExample struct {
	Manager *[]ObjectExampleManager `json:"manager"`
}

type ObjectExampleManager struct {
	Age  int32             `json:"age"`
	Name ObjectExampleName `json:"name"`
}

type ObjectExampleName struct {
	First string   `json:"first"`
	Last  []string `json:"last"`
}
