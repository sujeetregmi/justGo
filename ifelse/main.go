package main

import "fmt"

func main() {
	fmt.Println("IfElse in golang")

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular user."
	} else if loginCount > 10 {
		result = "Watch out."
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 2; num < 10 {
		fmt.Println("num less than 10.")
	} else {
		fmt.Println("num greater than 10.")
	}

	// if err !=nil{

	// }
}
