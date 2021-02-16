package main

import "fmt"

func main()  {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "BLue"

	fmt.Println(colors)

	fmt.Println(len(colors))


	var colos = []string{"red","green"}
	fmt.Println(colos)

	colos = append(colos, "purple")
	fmt.Println(colos)


	numbers := make([]string, 5, 5)
	numbers[0] = "Prakash"
	numbers[1] = "Sandeep"

	fmt.Println(numbers)
}