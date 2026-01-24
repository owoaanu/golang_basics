package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
	
	N := 10
	sum := 0
	
	for i:= 1; i<=N; i++ {
		sum += i
	}
	
	fmt.Println("Sum of first", N, "natural numbers is:", sum)
}
