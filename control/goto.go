package control

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GotoFunction() {
	reader := bufio.NewReader(os.Stdin)
	var input int

Loop:
	fmt.Print("Enter your age: ")
	inputString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inputString = strings.TrimSpace(inputString)

	age, err := strconv.Atoi(inputString)
	if err != nil {
		fmt.Println("Invalid input! Please enter a whole number.")
		goto Loop
	}

	input = age

	if input == 0 {
		fmt.Println("Invalid age! 🚫")
		goto Loop
	}

	if input < 18 {
		fmt.Println("You must be at least 18 years old! 👨🏼‍🦽")
		goto Loop
	}

	fmt.Println("You're eligible! 👍")
}
