package main

import "fmt"

func main() {
	a := [3]int{0, 0, 0}
	v := a[:]
	fmt.Printf("1 slice: %v len: %v cap: %v\n", v, len(v), cap(v))
	for i := 0; i <= 4; i++ {
		v = append(v, i)
		fmt.Printf("2 slice: %v len: %v cap: %v\n", v, len(v), cap(v))
	}
	fmt.Println(v)
}
