package main

import (
	"fmt"
)

// Старая функциональность: проверка четного числа
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

// Новая функциональность: калькулятор
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

// Новая задача: классификация числа
func classifyNumber() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	isPositive := number > 0
	isEven := number%2 == 0
	isDivisibleBy5 := number%5 == 0

	fmt.Printf("Положительное: %t\n", isPositive)
	fmt.Printf("Четное: %t\n", isEven)
	fmt.Printf("Делится на 5: %t\n", isDivisibleBy5)
}

func main() {
	// Выбор действия
	fmt.Println("Выберите действие:")
	fmt.Println("1 - Проверить четное число")
	fmt.Println("2 - Калькулятор")
	fmt.Println("3 - Классификация числа")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		checkEvenNumber()
	case 2:
		calculator()
	case 3:
		classifyNumber()
	default:
		fmt.Println("Неверный выбор")
	}
}
