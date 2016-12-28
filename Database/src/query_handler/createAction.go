package query_handler

import "fmt"

type CreateAction struct {
	keyword string
	minLength int
}

func (action *CreateAction) do(input [] string) {
	fmt.Println("table created from ", input)
}

func (action CreateAction) validate(input [] string) bool {
	correctFormat :=  len(input) > action.minLength && input[0] == action.keyword;
	return correctFormat
}

func (action CreateAction) execute (input [] string) {
	if (action.validate(input)) {
		action.do(input)
	}
}

