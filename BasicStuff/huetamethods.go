package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err) // Вывод: Ошибка: деление на ноль невозможно
		return
	}
	fmt.Println("Результат:", result)
}
