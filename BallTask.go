package main

import (
	"fmt"
)

func Ball() {
	// Input data
	var H, hMin, k float64
	fmt.Print("Enter the initial height (H): ")
	fmt.Scan(&H)
	fmt.Print("Enter the minimum height (h_min): ")
	fmt.Scan(&hMin)
	fmt.Print("Enter the bounce coefficient (k): ")
	fmt.Scan(&k)

	// Validate input data
	if H <= 0 || hMin <= 0 || hMin >= H || k <= 0 || k >= 1 {
		fmt.Println("Invalid input. Ensure that:")
		fmt.Println("1. H > 0")
		fmt.Println("2. h_min > 0 and h_min < H")
		fmt.Println("3. 0 < k < 1")
		return
	}

	// Count the number of bounces
	jumps := 0
	for H > hMin {
		H *= k
		jumps++
	}

	// Output the result
	fmt.Printf("Number of bounces: %d\n", jumps)
}
