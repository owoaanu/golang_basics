package main

import "fmt"

func main() {
	// slice are the most commonly used collection type
	// it is dynamic and it can grow
	// []type{...}

	results := []string{"pass", "fail", "pass", "pass"}

	fmt.Println("Results:", results)
	fmt.Println("results of first student: ", results[0])
	fmt.Println("result of the last student: ", results[len(results)-1])

	// creating an empty slice
	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30, 85)

	fmt.Println("Numbers:", nums)
}
