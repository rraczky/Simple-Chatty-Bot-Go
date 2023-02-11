package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)
	if number > 0 {
		fmt.Println("Positive!")
	} else if number == 0 {
		fmt.Println("Zero!")
	} else {
		fmt.Println("Negative!")
	}

	// Write your code here.
}
