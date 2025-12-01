package main

import (
	"learning-go/control"
	"learning-go/functions"
	"learning-go/helpers"
	"learning-go/types"
	"learning-go/variables"
)

func main() {
	helpers.DisplayOutputs([]helpers.ListOutput{
		{
			ModuleName: "Variables",
			Function:   variables.Main,
		},
		{
			ModuleName: "Types",
			Function:   types.Main,
		},
		{
			ModuleName: "Control",
			Function:   control.Main,
		},
		{
			ModuleName: "Functions",
			Function:   functions.Main,
		},
	})
}
