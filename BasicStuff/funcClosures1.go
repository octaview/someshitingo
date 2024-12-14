package main

import "fmt"

func createFilter(i int) func(k int) bool {
	threshold := i
	return func(k int) bool {
		if k > threshold {
			return true
		} else {
			return false
		}
	}
}

func main() {
	filter := createFilter(10)
	fmt.Println(filter(15))
	fmt.Println(filter(5))
}
