package functions

import "fmt"

/*
	🔧 Here, we create an anonymous function which acts as a function closure.

	💡 A function which has no name is called anonymous function.

	🌂 A closure is a function which refers reference variable from outside its body.

	💭 The function may access and assign to the referenced variables.
*/
func ClosureFunction(num uint) {
	squreNumber := func() uint {
		return num * num
	}

	fmt.Println("The square of", num, "is", squreNumber())
}
