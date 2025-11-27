package control

import (
	"learning-go/helpers"
)

func IfElseFunction() {
	const isSingle = true
	const isGay = true

	if isSingle && !isGay {
		helpers.Outputs([2]string{
			"Hey! You're single! 🕊️",
			"Do you have someone to love? 😁",
		})
	} else if isGay {
		helpers.Outputs([2]string{
			"Hey! You're gay! 😈",
			"Anyways, still gay! 👽",
		})
	} else {
		helpers.Outputs([2]string{
			"Oops! You're not single! 😢",
			"How many children do you have? 👼",
		})
	}
}

func SwitchCaseFunction() {
	const day = 3

	switch day {
	case 1:
		helpers.Outputs([2]string{
			"Today is Monday! 🍽️",
			"Time to get up! 🌅",
		})
	case 2:
		helpers.Outputs([2]string{
			"Today is Tuesday! 🍽️",
			"Time to get up! 🌅",
		})
	case 3:
		helpers.Outputs([2]string{
			"Today is Wednesday! 🍽️",
			"Time to get up! 🌅",
		})
	default:
		helpers.Outputs([2]string{
			"Today is another day! 🍽️",
			"Time to get up! 🌅",
		})
	}
}
