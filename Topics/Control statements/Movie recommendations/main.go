package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)

	// Code your switch or if...else-if statement here.
	switch {
	case age <= 14:
		fmt.Println("Toy Story 4")
	case age >= 15 && age <= 18:
		fmt.Println("The Matrix")
	case age >= 19 && age <= 25:
		fmt.Println("John Wick")
	case age >= 26 && age <= 35:
		fmt.Println("Constantine")
	case age > 35:
		fmt.Println("Speed")
	}
}
