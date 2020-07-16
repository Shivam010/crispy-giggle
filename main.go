package cg

import "fmt"

// Together is our new main.
func Together() string {
	fmt.Println("Together is our new main.")
	return Crispy() + "\n" + Giggle()
}

// Main calls the crispy giggles all the way together.
func Main() {
	fmt.Println(Together())
}
