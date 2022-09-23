package mapping

type Properties map[string]any

func (p *Properties) FillType() {
	for _, v := range *p {
		if filler, ok := v.(FillTyper); ok {
			filler.FillType()
		}
	}
}

type FillTyper interface {
	// FillType fills Type field if it is zero value.
	FillType()
}

type onScriptError string

const (
	Continue onScriptError = "continue"
	Fail     onScriptError = "fail"
)
