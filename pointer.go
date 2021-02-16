package main


import "fmt"

func main()  {
	//Type of ptr is pointer of integer type
	var ptr *int
	c := 20
	ptr = &c
	*ptr = 10
	fmt.Println(*ptr)
}