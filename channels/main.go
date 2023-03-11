// channels uses <- operator to send and recieve data.
// ch <- v  sends V to channel ch
// v := <-ch receives data from ch and assign it to v

package main

import "fmt"

func main() {
	fmt.Println("Channels in Go lang.")
	a := []int{5, 4, 5, 3, 5, 3, 32, 1}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}
