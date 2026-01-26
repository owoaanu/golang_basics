package main

import "fmt"

func main(){
	views := []int{10,20,45,50,60}
	
	// for range is the common approach to loop over slices in Go
	
	total := 0
	for i,v := range views{
		fmt.Println("day",i, "views",v)
		total += v
	}
	
	fmt.Println("Total views:", total)
}