package functions

/*
	🌀 In Go programming, calling same function from within the function is known as recursion.

	😃 It is always a good idea to break a problem into multiple tasks.

	🥳 Let us see a program to calculate factorial value in Go programming using recursion.
*/
func FactorialRecursion(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * FactorialRecursion(num-1)
}
