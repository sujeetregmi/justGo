package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang
	regmi := User{"Sujeet", "regmi@gmail.com", true, 22}
	fmt.Println(regmi)
	fmt.Printf("This is Detail of user: %+v", regmi)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
