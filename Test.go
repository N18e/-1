package main

import (
	"fmt"
)

// Стара функціональність
func checkEvenNumber() {
	var number int
	fmt.Println("Введите Ваше число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Printf("Число %d четное: true\n", number)
	} else {
		fmt.Printf("Число %d четное: false\n", number)
	}
}

// Новий калькулятор
func calculator() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
		} else {
			fmt.Printf("Результат: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Ошибка: недопустимый оператор")
	}
}

func main() {
	// Вибір функціональності
	fmt.Println("Выберите действие:")
	fmt.Println("1 - Проверить четное число")
	fmt.Println("2 - Калькулятор")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		checkEvenNumber()
	} else if choice == 2 {
		calculator()
	} else {
		fmt.Println("Неверный выбор")
	}
}
