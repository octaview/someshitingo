package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 0
	f0 := 0
	f1 := 1
	return func() int {
		if count == 0 {
			count++
			return f0
		} else if count == 1 {
			count++
			return f1
		} else {
			next := f0 + f1
			f0, f1 = f1, next
			count++
			return next
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
