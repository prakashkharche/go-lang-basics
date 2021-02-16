package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakMutatingSound() {
	d.Sound = d.Sound + d.Sound + d.Sound
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakMutatingSoundPtr() {
	d.Sound = d.Sound + d.Sound + d.Sound
	fmt.Println(d.Sound)
}

func main()  {
	dexter := Dog{"Cute", 10, "Bow bow"}
	fmt.Println(dexter)

	fmt.Println(dexter.Breed)
	dexter.Speak()

	dexter.Sound = "Ewwww"

	dexter.Speak()


	dexter.SpeakMutatingSound()
	dexter.Speak()

	dexter.SpeakMutatingSoundPtr()
	dexter.Speak()
}