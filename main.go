package main

import (
	"cycloid-challenge-words/dic"
	"fmt"
)

///
func main() {

	fmt.Println("START")
	first := "cat"
	last := "dog"

	if len(first) != len(last) {
		fmt.Println("Error: Incorrect Inputs")
		return
	}
	words := dic.Dictionary(len(first))
	fmt.Println(words)

}
