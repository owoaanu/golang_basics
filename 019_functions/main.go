package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sumAndProduct (a int, b int) (int, int) { // for multiple return types
	sum := a+b
	product:= a*b
	
	return sum, product
}
func main() {
	sum1 := add(10, 20)
	s, p := sumAndProduct(10, 20)
	fmt.Println("Sum:", s)
	fmt.Println("Product:", p)
	
	fmt.Println("Sum1:", sum1)
	
	// sometimes, when you don't want to use the return value, use _ for a placeholder
	onlySum, _ := sumAndProduct(15, 25)
	fmt.Println("Only Sum:", onlySum)
}
