package main

import "fmt"

func main()  {

	//make vs new  - make allocated memory + initializes, 
	m := make(map[string]string)

	m["k1"] = "v1"
	m["k2"] = "v2"
	m["k3"] = "v3"

	fmt.Println(m, m["k2"])

	for k,v := range m {
		fmt.Println("Key ", k, "Value " ,v)
	}
}