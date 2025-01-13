package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите Ваше число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Printf("Число %d четное: true\n", number)
	} else {
		fmt.Printf("Число %d четное: false\n", number)
	}
}
