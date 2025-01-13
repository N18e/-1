package main

import (
	"fmt"
)

func main() {
	s := "gopher"
	var name string = "Sergey"
	fmt.Println("Hello and welcome, %s!", s)
	fmt.Println("Hi my name is", name)
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
