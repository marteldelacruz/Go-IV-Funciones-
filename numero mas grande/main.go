package main

import "fmt"

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

func main() {
	fmt.Println(largestNumber(12, 44, 1, 2, 3, 77, 111111))
}
