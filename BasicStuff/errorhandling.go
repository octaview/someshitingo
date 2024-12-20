package main

import (
	"errors"
	"fmt"
)

func main() {
	var s int
	fmt.Scan(&s)
	message, error := canIPass(s)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(message)
}

func canIPass(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Проходите", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понравится эта музыка?", nil
	} else if age >= 65 {
		return "Это уже слишком ", errors.New("you're too old")
	}
	return "Контроль возраста не пройден", errors.New("you're too young")
}
