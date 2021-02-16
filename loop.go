package main

import "fmt"

func main()  {
	colors := []string{"red", "green", "blue"}

	for i:=0; i<3; i++ {
		fmt.Println("Conv loop", colors[i])
	}


	for i:=0; i<len(colors); i++ {
		fmt.Println("Len loop", colors[i])
	}

	for i := range(colors) {
		fmt.Println("Range loop", colors[i])
	}


	//Something like WHILE loop

	sum :=1

	for sum < 10000 {
		sum += 2
		if (sum > 200) {
			goto reached200
		}
	}

	fmt.Println(sum)

	reached200: {
		fmt.Println("Reached 200")
	}
 }