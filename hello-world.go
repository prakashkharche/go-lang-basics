package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World..!")
	fmt.Println(strings.ToUpper("this is going to be loud."))
	fmt.Println(strings.Title("this is going to be loud."))

	fmt.Println("Equals?", strings.EqualFold("Prakash", "prakash"))

	str1 := "Prakash"
	str2 := "Kharche"
	str3 := "Rocks"
	str4 := 22.4
	length,err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println(length)
		fmt.Println(float64(str4))
	}

	var s string
	fmt.Scanln(&s)
	fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	fmt.Println(str)

	fl,_ := strconv.ParseFloat("32.09", 64)
	fmt.Println(fl)
}