package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println("List of languages:", languages)

	// deleting in golang using delete keyword
	delete(languages, "py")
	fmt.Println(languages)

	// control flow in map (for loop)
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v", key, value)
	}
}
