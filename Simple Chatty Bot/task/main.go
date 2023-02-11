package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")
	fmt.Println("Please, remind me your name.")

	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	var rem3, rem5, rem7, age int
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
	fmt.Println("Now I will prove to you that I can count to any number you want.")

	// read a number and count to it here
	var count int
	fmt.Scan(&count)
	for i := 0; i <= count; i++ {
		fmt.Println(i, "!")
	}

	fmt.Println("Completed, have a nice day!")
}
