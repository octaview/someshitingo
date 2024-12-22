package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin(4, 2, 5, 1))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
	}
	return min
}
