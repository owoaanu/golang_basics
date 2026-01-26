package main

import "fmt"

func main(){
	// map[keyTpe]valueType
	
	ages := map[string]int{
		"sam":25,
		"johnson":30,
		"anna":22,
	}
	
	fmt.Println("Ages map:", ages)
	fmt.Println("Anna's age:", ages["anna"])
	fmt.Println(len(ages))
	
	
	//make(map[k]v)  to create an empty map using make function
	
	scores := make(map[string]int)
	
	scores["a"] = 90
	fmt.Println("Scores map:", scores, scores["a"])
	
	users := map[string]string{
		"u1": "sam",
		"u2": "john",
		"u3": "anna",
	}
	users["u4"] = "Johan"
	
	fmt.Println("Users map:", users, users["u2"])
	
	delete(users, "u3") // how to delete elements from mapsf
	
	fmt.Println("Users map after deletion:", users)
	
}