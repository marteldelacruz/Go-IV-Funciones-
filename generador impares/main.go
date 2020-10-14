package main

import "fmt"

func generateOdds() func() uint {
	i := uint(0)
	return func() uint {
		var par = i
		i += 2
		return par + 1
	}
}

func main() {
	nextOdd := generateOdds()

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
