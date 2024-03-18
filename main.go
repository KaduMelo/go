package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Walk() {
	fmt.Println(p.name + " is walking")
}

func main() {
	var a int
	a = 10
	fmt.Println(a)

	// array
	// var b [5]int

	// slice
	// var c []int

	var pessoa Person
	pessoa.name = "Chessi"
	pessoa.age = 32

	fmt.Println(pessoa.name + " is " + string(pessoa.age) + " years old")

	pessoa.Walk()
}
