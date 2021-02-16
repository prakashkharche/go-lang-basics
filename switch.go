package main

import "fmt"

func main()  {
	
	s := "prakash"

	switch s {
	case "prakash":
		fmt.Println("Cool Prakash")
	case "kharche":
		fmt.Println("Cool Kharche")
	default:
		fmt.Println("No cool")
	}


	v := 10
	switch {
	case v>0:
		fmt.Println("Value > 0")
		//fallthrough - To work as previous
	case v==0:
		fmt.Println("Value = 0")
	case v<0:
		fmt.Println("Value < 0")
	}
}