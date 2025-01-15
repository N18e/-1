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

// classify a number
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

// Print numbers in a range, skipping multiples of 3
func printRangeSkippingMultiplesOf3() {
	var start, end int
	fmt.Print("Enter the start of the range: ")
	fmt.Scan(&start)
	fmt.Print("Enter the end of the range: ")
	fmt.Scan(&end)

	if start > end {
		fmt.Println("Error: Start of the range is greater than the end.")
		return
	}

	fmt.Println("Numbers in the range, skipping multiples of 3:")
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Calculate the sum and product of numbers from 1 to N
func calculateSumAndProduct() {
	var n int
	fmt.Print("Enter a positive integer N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Error: Number must be positive.")
		return
	}

	sum := 0
	product := 1
	for i := 1; i <= n; i++ {
		sum += i
		product *= i
	}

	fmt.Printf("Sum of numbers from 1 to %d: %d\n", n, sum)
	fmt.Printf("Product of numbers from 1 to %d: %d\n", n, product)
}

func main() {
	// Menu selection
	fmt.Println("Select an action:")
	fmt.Println("1 - Check if a number is even")
	fmt.Println("2 - Calculator")
	fmt.Println("3 - Classify a number")
	fmt.Println("4 - Print numbers in a range, skipping multiples of 3")
	fmt.Println("5 - Calculate the sum and product of numbers")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		checkEvenNumber()
	case 2:
		calculator()
	case 3:
		classifyNumber()
	case 4:
		printRangeSkippingMultiplesOf3()
	case 5:
		calculateSumAndProduct()
	default:
		fmt.Println("Invalid choice")
	}
}
