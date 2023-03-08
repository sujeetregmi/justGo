package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang.")
	content := "This is content of file."
	file, err := os.Create("./myfiles.txt")

	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is:", length)
	defer file.Close()
	readFile("./myfiles.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	// fmt.Println("Text data inside the file is:\n",databyte)
	fmt.Println("Text data inside the file is:\n", string(databyte))
}

func checkNilError(err error){
	if err!=nil{
		panic(err)
	}
}