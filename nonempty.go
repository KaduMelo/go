package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data))
	fmt.Println(data)
}

// nonempty is a example of the algorithm in-place to slices
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}
