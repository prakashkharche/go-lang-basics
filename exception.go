package main

import (
	"fmt"
	"os"
	"errors"
)
func main() {
	f, error := os.Open("file_is_absent.txt")

	if error == nil {
		fmt.Println("File read successfully", f)
	} else {
		fmt.Println(error.Error())
	}

	myError := errors.New("A custom error message")
	fmt.Println(myError)

	attendace := map[string]bool {
		"Prakash":true,
		"Kharche":false}

	//ok is presence check inside the map
	if attended, ok := attendace["Prakash"]; ok {
		fmt.Println(attended)
	}

	if attended, ok := attendace["Prakash1"]; ok {
		fmt.Println(attended)
	} else {
		fmt.Println("No info")
	}
}