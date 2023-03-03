package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang.")
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[4] = "Cherry"

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var veg =[3]string {"Potato","Brinjal","Cauliflower"}
	fmt.Println(veg)
}
