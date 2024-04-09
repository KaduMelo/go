package main

import (
	"fmt"
)

func main() {
	// this is an array
	a := [6]string{"this", "is", "a", "collection", "the", "words"}

	// here creates a slice from an array
	s := a[:]

	fmt.Printf("value a[3]: %q\n", a[3])
	fmt.Printf("value s[3]: %q\n", s[3])

	// like s is a pointer to a, when changing s too change a
	fmt.Println(`— Like "s" pointer to "a", change "s" to change "a"`)
	s[3] = "array"
	fmt.Printf("value a[3]: %q\n", a[3])
	fmt.Printf("value s[3]: %q\n", s[3])

	// lets go now do an append in slice
	fmt.Println(`— Do append in "s" create a new array`)
	s = append(s, "!")

	// using append it was created a new array copying the data
	// now lets change the value s[3] to show this
	s[3] = "slice"
	fmt.Printf("value a[3]: %q\n", a[3])
	fmt.Printf("value s[3]: %q\n", s[3])

	// how were the arrays
	fmt.Printf("value a: %q\n", a)
	fmt.Printf("value s: %q\n", s)
}
