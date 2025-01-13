package main

import (
	"fmt"
)

// Old functionality: check for even number
func checkEvenNumber() {
	var number int
	fmt.Println("Enter your number: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Printf("Number %d is even: true\n", number)
	} else {
		fmt.Printf("Number %d is even: false\n", number)
	}
}

// New functionality: calculator
func calculator() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: division by zero")
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Error: invalid operator")
	}
}

// New task: classify a number
func classifyNumber() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	isPositive := number > 0
	isEven := number%2 == 0
	isDivisibleBy5 := number%5 == 0

	fmt.Printf("Positive: %t\n", isPositive)
	fmt.Printf("Even: %t\n", isEven)
	fmt.Printf("Divisible by 5: %t\n", isDivisibleBy5)
}

func main() {
	// Select an action
	fmt.Println("Select an action:")
	fmt.Println("1 - Check if the number is even")
	fmt.Println("2 - Calculator")
	fmt.Println("3 - Classification of number")
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
		fmt.Println("Wrong choice")
	}
}
