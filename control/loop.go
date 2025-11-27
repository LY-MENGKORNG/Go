package control

import (
	"fmt"
)

func ForLoop() {
	fmt.Println("For Loop Function")
	for i := 0; i < 5; i++ {
		var icons string = ""
		for j := 0; j <= i; j++ {
			icons += "👋"
		}
		fmt.Printf("Hi mom! %s\n", icons)
	}
	fmt.Println("")
}

func ForRangeLoop() {
	fmt.Println("For Range Loop Function")
	for i := range 5 {
		var icons string = ""

		for range i + 1 {
			icons += "👋"
		}

		fmt.Printf("Hi mom! %s\n", icons)
	}
	fmt.Println("")
}

func BreakLoop() {
	fmt.Println("Break Loop Function")
	for i := range 5 {
		var icons string = ""
		for j := 0; j <= i; j++ {
			icons += "👋"
		}
		fmt.Printf("Hi mom! %s\n", icons)
		if i == 2 {
			fmt.Println("I'm tired of this! 😴. Let's go home! 🏠")
			break
		}
	}
	fmt.Println("")
}

func ContinueLoop() {
	fmt.Println("Continue Loop Function")
	for i := range 5 {
		var icons string = ""
		for j := 0; j <= i; j++ {
			icons += "👋"
		}
		if i == 2 {
			fmt.Println("I'm tired of this! 😴. Let's take a break! 🏖️")
			continue
		}
		fmt.Printf("Hi mom! %s\n", icons)
	}
	fmt.Println("")
}
