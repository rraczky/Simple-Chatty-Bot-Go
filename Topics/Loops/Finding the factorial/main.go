package main

import "fmt"

func main() {
	var number int
	var factorial int = 1
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println(factorial)

}
