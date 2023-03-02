package main

import "fmt"

const LogInToken string = "sdfsdgdsf" //public

func main() {
	var username string = "sujeet regmi"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable is of type %T \n", isLoggedIn)

	var smallValue uint = 255
	fmt.Println(smallValue)
	fmt.Printf("The variable is of type %T \n", smallValue)

	var smallFloat float64 = 255.534634663463624365252526732
	fmt.Println(smallFloat)
	fmt.Printf("The variable is of type %T \n", smallFloat)

	// default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("This is of type: %T \n ", anotherVariable)

	// implicit type

	var website = "www.github.com"
	fmt.Println(website)
	fmt.Printf("%T \n", website)

	// no var style
	numberOfuser := 2323443
	fmt.Println(numberOfuser)
}
