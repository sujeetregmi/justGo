package main

import "fmt"

func main() {
	fmt.Println("Abour Pointers")
	// basic pointer working in golang
	number := 33
	var ptr = &number

	fmt.Println("Pointer address is :", ptr)
	fmt.Println("Value of the pointer stored in address is", *ptr) //actual value

	*ptr = *ptr + 2
	fmt.Println("New Value is:",number)
}
