package main

import "fmt"

func fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return fibonacci(number-2) + fibonacci(number-1)
}

func largestNumber(args ...int) int {
	// this will store the number
	largest := args[0]

	for i, v := range args {
		if largest < args[i] {
			largest = v
		}
	}
	return largest
}

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
