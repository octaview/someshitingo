package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	results := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		if len(s) > 10 {
			results[i] = fmt.Sprintf("%s%d%s", string(s[0]), len(s)-2, string(s[len(s)-1]))
		} else {
			results[i] = s
		}
	}
	for _, value := range results {
		fmt.Println(value)
	}
}
