package main

import "fmt"

// variadic funtions allow the creation of a function that can take
// a variable number of parameters to the funstion, iff the params
// are all of the same types
func sumAll(nums ...int) int {
	total := 0

	for _, currentValue := range nums {
		total += currentValue
	}

	return total

}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6))

	// it also works with slices
	scores := []int{120, 23, 33}

	fmt.Println("from slices", sumAll(scores...))

	// anoymous function
	res := func (a int)int{
		return a*2
	}
	
	fmt.Println("output from the use of the anonymous function",res(5))
	
	
	//IIFE - immediately invoked function expression
	var res1 = func(a int, b int) int{
	return (a+b)*3
	}(5,6)
	
	fmt.Println("output from the IIFE functions", res1)
}
