package main

import "fmt"

func print_hello() {
	fmt.Printf("Hello!\n")
}

func main() {
	// defer
	defer print_hello()

	// if, else if, else
	age := 19

	if age < 18 {
		fmt.Printf("Not old enough!\n")
	} else if age > 60 {
		fmt.Printf("Too old!\n")
	} else {
		fmt.Printf("Welcome!\n")
	}
}
