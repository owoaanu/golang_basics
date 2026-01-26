package main

import (
	"fmt"
)

func main() {
	// arrays are fixed and cannot grow - very very important
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 50

	fmt.Println("Marks:", marks)
	
	// array literal - a quick way of creating a fixed array
	var scores = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Scores:", scores)
	
	fmt.Println("Size of the scores array: ",len(scores))
	
	
}
