package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Golang")

	// creating slices in golang
	var fruits = []string{"apple", "banana", "cherry"}
	println(len(fruits))
	fmt.Printf("data type of fruits is %T \n", fruits)

	fruits = append(fruits, "Pineapple", "Guava", "jackfruit")
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	// memory management in golang using make or new keyword

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 936
	highScore[2] = 239
	highScore[3] = 254

	fmt.Println(highScore)

	highScore = append(highScore, 555)
	fmt.Println(highScore)
	fmt.Printf("type of highscore is %T", highScore)
	fmt.Println(len(highScore))
	// sort method
	sort.Ints(highScore)
	fmt.Println(highScore)

	// how to remove the value from slices based on index
	var courses = []string{"nodejs", "javascript", "swift", "django", "golang"}
	fmt.Println(courses)
	var index int = 2

	// deleting swift from the courses
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
