package helpers

import (
	"fmt"
)

type ListOutput struct {
	ModuleName string
	Function   func()
}

func DisplayOutputs(outputList []ListOutput) {
	for _, output := range outputList {
		fmt.Println("")
		fmt.Printf("🌳🌳🌳 %v 🌳🌳🌳 \n", output.ModuleName)
		output.Function()
	}
}

func Outputs(strings [2]string) {
	for _, s := range strings {
		fmt.Println(s)
	}
	fmt.Print("\n")
}
