package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordcount := make(map[string]int)
	for _, value := range words {
		_, ok := wordcount[value]
		if !(ok) {
			for i := 0; i < len(words); i++ {
				if words[i] == value {
					wordcount[value] += 1
				}
			}
		}
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}
