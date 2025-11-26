package variables

import (
	"learning-go/helpers"
)

var packageVariable string = "Hey! I'm a package variable. anyone can access me. 🫡"
var GlobalVariable string = "Hey! I'm a global variable. I can be accessed from any function in the program. 🚀"

func ScopingVariables() {
	var localVariable string = "Hey! I'm a local variable. I can only be accessed within this function. 👻"

	helpers.Outputs([]string{
		"🚨 Local variables are defined inside a function and can only be accessed within that function.",
		"	💡 " + localVariable,
	})

	helpers.Outputs([]string{
		"📦 Package variables are defined in the package scope and can be accessed from any function in the package.",
		"	💡 " + packageVariable,
	})

	helpers.Outputs([]string{
		"🌎 Global variables are defined outside a function and can be accessed from any function in the program.",
		"	💡 " + GlobalVariable,
	})
}
