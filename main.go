package cg

import "fmt"

// Together is our new main.
func Together() string {
	fmt.Println("Together is our new main.")
	return Crispy() + "\n" + Giggle()
}

// Main now, calls the crispy giggles all the way together with a new level of attendency.
func Main(opt ...interface{}) {
	opt = append(opt, Together())
	fmt.Println(opt...)
}
