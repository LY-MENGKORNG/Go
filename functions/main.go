package functions

import "fmt"

/*
	 	ⓕ In Go, functions are the basic building blocks. A function is used to break a large problem into smaller tasks.
		We can invoke a function several times, hence functions promote code reusability. There are 3 types of functions in Go:

		🪄 Normal functions with an identifier
		🪄 Anonymous or lambda functions
		🪄 Method (A function with a receiver)
*/
func Main() {
	var number uint = 9

	NormalFunction()

	var factorial uint = FactorialRecursion(number)
	fmt.Println("Factorial of", number, "is", factorial)

	ClosureFunction(number)

}
