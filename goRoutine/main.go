package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Routines in Go.")
	go say("World") //create a new go routine
	say("Hello")    //current go routine
}
