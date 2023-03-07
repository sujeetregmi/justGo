package main

import "fmt"

func main() {
	fmt.Println("Functions in golang.")
	result:=adder(10,13)
	fmt.Println("The sum is:",result)
	proResult,proMessage := proAdder(2,3,4,5,3,5,3,345,5,3,3,3,3,4,43,5,3,535,3,535,3,5)
	fmt.Println("Proresult is:",proResult)
	fmt.Println("Pro message is:",proMessage)
}

func adder(a int, b int) int{
	return a+b
}

func proAdder(values... int) (int,string){
	total := 0

	for _, val := range values{
		total +=val
	}
	return total,"This is Pro message."
}