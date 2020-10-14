package main

import "fmt"

func intercambia(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a, b int

	// ask for values
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	// interchange
	intercambia(&a, &b)

	// print final values
	println("a: ", a, "\t b: ", b)
}
