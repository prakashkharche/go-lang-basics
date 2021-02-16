package main

import "fmt"

func main() {

	if v := 10; v > 0 {
		fmt.Println("Greater than 0")
	} else if v < 0 {
		fmt.Println("Less than 0")
	} else {
		fmt.Println("Its 0")
	}
}