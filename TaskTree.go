package main

import (
	"fmt"
)

// MultiplicationTable
func MultiplicationTable() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

	// Sum of Even Numbers
	fmt.Print("\nEnter a number N: ")
	var n int
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of even numbers: %d\n", sum)

	// Countdown
	fmt.Print("\nEnter a positive integer: ")
	var countdown int
	fmt.Scan(&countdown)
	for countdown >= 0 {
		fmt.Println(countdown)
		countdown--
	}
	fmt.Println("Liftoff!")
}
