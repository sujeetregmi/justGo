// select is used to manage multiple channels at once

package main

import "fmt"

func main() {
	fmt.Println("Select in golang.")
	c :=make(chan int)
	quit := make(chan int)
	go func () {
		for i:=0;i<10;i++{
			fmt.Println(<-c)
		}
		quit <-0
	}()
	fib(c,quit)
}

func fib(c, quit chan int){
	x,y:=1,1
	for{
		select{
		case c <-x:
			x,y = y,x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}
