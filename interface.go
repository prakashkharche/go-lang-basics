package main

import "fmt"

// interface{} -- this satisfies every type in golang, It says interface with no/empty method.

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Bow Bow"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow"
}

type Cow struct {
}

func (c Cow) Speak() string {
	return "Mooooooo"
}
func main() {
	dexter := Dog{}
	animal := Animal(dexter) // This is the way we cast
	fmt.Println(animal.Speak())

	animals := []Animal{Dog{}, Cat{}, Cow{}}

	for i := range(animals) {
		fmt.Println(animals[i].Speak())
	}
}