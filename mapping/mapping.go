package mapping

type Properties map[string]Params

type Params struct {
}

type onScriptError string

const (
	Continue onScriptError = "continue"
	Fail     onScriptError = "fail"
)
