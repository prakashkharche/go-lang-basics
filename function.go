package main

import "fmt"


func main()  {
	fmt.Println("Sum : ",sum(2,3))
	fmt.Println("Sum all values: ",sum2(2,3,4,5,6))

	sum, count := sumAndCount(1,2,3,4,5,6,7)
	fmt.Println("Sum :", sum, " Count:", count)
}

func sum(val1 int, val2 int) int {
	return val1+val2
}

func sum2(values ...int) int {
	sum := 0
	for i:= range(values) {
		sum += values[i]
	}
	return sum
}


//Order of return has to be kept in mind here.
func SumAndCount(values ...int) (int, int) {
	sum :=0
	count :=0
	for i:= range(values) {
		sum += values[i]
		count++
	}
	return sum, count
}


// no need to maintain order here.
func SumAndCountNakedReturn(values ...int) (sum int, count int) {
	sum =0
	count =0
	for i:= range(values) {
		sum += values[i]
		count++
	}
	return count, sum
}