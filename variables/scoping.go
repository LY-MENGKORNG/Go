package variables

import (
	"learning-go/helpers"
)

var packageVariable string = "Hey! I'm a package variable. anyone can access me within the package. 🫡"
var GlobalVariable string = "Hey! I'm a global variable. I can be accessed from any function in the program. 🚀"

func ScopingVariables() {
	var localVariable string = "Hey! I'm a local variable. I can only be accessed within this function. 👻"

	helpers.Outputs([2]string{
		"🚨 Local variables are defined inside a function and can only be accessed within that function.",
		"|_ " + localVariable,
	})

	helpers.Outputs([2]string{
		"📦 Package variables are defined in the package scope and can be accessed from any function in the package.",
		"|_ " + packageVariable,
	})

	helpers.Outputs([2]string{
		"🌎 Global variables are defined outside a function and can be accessed from any function in the program.",
		"|_ " + GlobalVariable,
	})
}
