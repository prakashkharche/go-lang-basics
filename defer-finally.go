package main

import "fmt"

func main() {
	fmt.Println("statement1")
	defer fmt.Println("statement2")
	defer fmt.Println("statement3")
	defer fmt.Println("statement4")
	defer fmt.Println("statement5")
}