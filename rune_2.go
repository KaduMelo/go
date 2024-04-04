package main

import "fmt"

func main() {

	slc := []rune{'♛', '♠', '♧', '♡', '♬'}

	for i, value := range slc {
		fmt.Printf("\nCharacter: %c, Unicode: %U, Position: %d\n", value, value, i)
	}
}
