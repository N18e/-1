// Main program with task selection
package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Main Menu:")
		fmt.Println("1. Check if a number is even")
		fmt.Println("2. Basic calculator")
		fmt.Println("3. Classify a number")
		fmt.Println("4. Print numbers in a range (skip multiples of 3)")
		fmt.Println("5. Calculate sum and product up to N")
		fmt.Println("6. Multiplication table")
		fmt.Println("7. Ball bounce counter")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 0 {
			fmt.Println("Exiting the program. Goodbye!")
			break
		}

		switch choice {
		case 1:
			checkEvenNumber()
		case 2:
			basicCalculator()
		case 3:
			classifyNumber()
		case 4:
			printRangeSkippingMultiplesOf3()
		case 5:
			calculateSumAndProduct()
		case 6:
			multiplicationTable()
		case 7:
			ballBounceCounter()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println() // Add a blank line for readability
	}
}

// Task 1: Even number check
func checkEvenNumber() {
	var num int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&num)
	isEven := num%2 == 0
	fmt.Printf("The number %d is even: %t\n", num, isEven)
}

// Task 2: Basic calculator with error handling
func basicCalculator() {
	var num1, num2 float64
	var operator string
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter an operator (+, -, *, /): ")
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
			fmt.Println("Error: Division by zero")
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Error: Invalid operator")
	}
}

// Task 3: Classifying a number
func classifyNumber() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	isPositive := num > 0
	isEven := num%2 == 0
	isDivisibleBy5 := num%5 == 0
	fmt.Printf("Positive: %t\nEven: %t\nDivisible by 5: %t\n", isPositive, isEven, isDivisibleBy5)
}

// Task 4: Printing numbers in a range (skip multiples of 3)
func printRangeSkippingMultiplesOf3() {
	var start, end int
	fmt.Print("Enter the start of the range: ")
	fmt.Scan(&start)
	fmt.Print("Enter the end of the range: ")
	fmt.Scan(&end)
	fmt.Println("Numbers in the range not divisible by 3:")
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Task 5: Calculating sum and product of numbers up to N
func calculateSumAndProduct() {
	var n int
	fmt.Print("Enter a positive integer N: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Error: The number must be positive")
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

// Task 6: Multiplication table
func multiplicationTable() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Printf("Multiplication table for %d:\n", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}

// Task 7: Ball bounce counter
func ballBounceCounter() {
	var H, h_min, k float64
	fmt.Print("Enter the initial height (H): ")
	fmt.Scan(&H)
	fmt.Print("Enter the minimum height (h_min): ")
	fmt.Scan(&h_min)
	fmt.Print("Enter the bounce coefficient (k): ")
	fmt.Scan(&k)

	if H <= 0 || h_min <= 0 || h_min >= H || k <= 0 || k >= 1 {
		fmt.Println("Error: Invalid input values. Make sure H > 0, h_min > 0, h_min < H, and 0 < k < 1.")
		return
	}

	bounces := 0
	for H > h_min {
		H *= k
		bounces++
	}
	fmt.Printf("The ball bounces %d times before falling below the minimum height.\n", bounces)
}
