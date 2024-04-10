package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 3))
}

// func remove(s []int, index int) []int {
// 	p := make([]int, len(s))
// 	for i := 0; i < len(s); i++ {
// 		if i != index {
// 			p = append(p, s[i])
// 		}
// 	}
// 	return p
// }

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
