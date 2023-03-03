package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza")

	// comma ok syntax or error ok syntax
	// there is no try catch in go so we can use error ok syntax instead try catch

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating :", input)

}
