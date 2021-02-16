package main

import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Weekday())

	tomorrow := now.AddDate(0,0,1)
	fmt.Println(tomorrow.Month())
	fmt.Println(tomorrow.Day())
	fmt.Println(tomorrow.Weekday())
}