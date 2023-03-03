package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handelling time in golang.")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// always use the format like passed below line with the same exact way. No other way to format exact date.
	fmt.Println(presentTime.Format("01-02-2006 15:04 Monday"))

	createdDate := time.Date(2023, time.August, 10, 23, 23, 22, 0, time.UTC)
	fmt.Println(createdDate)
	// format created date as formated in present time.
}
