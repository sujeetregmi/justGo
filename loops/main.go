package main

import "fmt"

func main() {
	fmt.Println("Loops in golang.")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for loop in golang
	// for d:=1; d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for index, days := range days {
		fmt.Println(index, days)
	}

	rougeValue := 1

	for rougeValue < 10 {
		// if rougeValue == 5 {
		// 	break
		// }

		if rougeValue==2{
			goto lco
		}
		if rougeValue ==6{
			rougeValue++
			continue
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

	lco:
		fmt.Println("Jumping at learncodeonline.in")
}
