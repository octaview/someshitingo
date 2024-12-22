package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := make([]string, 1)
	messages[0] = "1"
	fmt.Println(messages)
}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}
	return nil
}
