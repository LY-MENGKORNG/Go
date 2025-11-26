package main

import (
	"learning-go/helpers"
	"learning-go/types"
	"learning-go/variables"
)

func main() {
	helpers.DisplayOutputs([]func(){
		variables.Main,
		types.Main,
	})
}
