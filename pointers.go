package main

import "fmt"

func main() {

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	// Tests to function f()
	z, v := f()
	fmt.Println(z, v)
	fmt.Println(*z, *v)
	fmt.Println(z == v)

	// Tests to function incr()
	i := 1
	incr(&i)
	fmt.Println(i)
}

func f() (*int, *int) {
	v := 1
	z := 1
	return &v, &z
}

func incr(p *int) int {
	*p++ // increment the value to what p aim
	return *p
}
