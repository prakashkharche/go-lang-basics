package main

import (
	"fmt"
	"math/big"
)


func main()  {
	i1, i2 ,i3 := 1,2,3
	fmt.Println(i1+i2+i3)

	f1,f2,f3 := 12.4, 12.5, 9.87
	fmt.Println(f1+f2+f3)

	var b1,b2,b3, bigSum big.Float

	b1.SetFloat64(10.9)
	b2.SetFloat64(12.9)
	b3.SetFloat64(2.9)

	bigSum.Add(b1, b2).Add(bigSum, b3)

	fmt.Println(bigSum)
}