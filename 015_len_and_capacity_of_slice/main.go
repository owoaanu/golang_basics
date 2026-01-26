package main

import(
	"fmt"
)
func main() {
	// len is how many elements are in the slice
	// capacity is how many elements you can stor
	
	// make([]T, len, cap)
	scores:= make([]int, 0, 5)
	
	fmt.Println("scores:", scores)
	fmt.Println("len:", len(scores))
	fmt.Println("cap:", cap(scores))
	
	scores = append(scores, 10)
	scores = append(scores, 20, 90)
	scores = append(scores, 30,855)
	
	fmt.Println("scores:", scores)
	fmt.Println("len:", len(scores))
	fmt.Println("cap:", cap(scores))
	
	// if the capacity is exceeded, a new underlying slice is created, 
	// usually it doubles the capacity if the initial slice
	fmt.Println("Appending one more element to exceed capacity")
	scores = append(scores, 100)
	
	fmt.Println("scores:", scores)
	fmt.Println("len:", len(scores))
	fmt.Println("cap:", cap(scores))
	
	todos := []string{"read", "workout"}
	
	more := []string{"learn golang"}
	
	// appending a slice into another using ...
	todos = append(todos, more...)
	
	fmt.Println("todos:", todos)
}
