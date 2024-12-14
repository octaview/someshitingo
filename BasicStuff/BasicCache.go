package main

import (
	"fmt"
)

func createCache() func(i int) int {
	newCache := make(map[int]int)
	return func(i int) int {
		_, ok := newCache[i]
		if ok {
			return newCache[i]
		} else {
			newCache[i] = i * i
			return newCache[i]
		}
	}
}

func main() {
	cache := createCache()
	fmt.Println(cache(4))
	fmt.Println(cache(4))
}
