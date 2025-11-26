package helpers

import (
	"fmt"
)

func DisplayOutputs(function []func()) {
	for _, f := range function {
		fmt.Println("------------------")
		f()
	}
}

func Outputs(strings []string) {
	for _, s := range strings {
		fmt.Println(s)
	}
	fmt.Print("\n")
}
