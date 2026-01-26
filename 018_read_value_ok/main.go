package main

import "fmt"

func main(){
	points := map[string]int{
		"a":10,
		"b":0, // valid value
		"c":30,
	}
	
	fmt.Println("a:", points["a"])
	fmt.Println("b:", points["b"])
	
	fmt.Println("d:", points["d"]) // not existing key, returns zero value of int type which is 0
	
	
	valC, okC := points["c"]
	fmt.Println("c:", valC, "exists?", okC)
	
	valD, okD := points["d"]
	fmt.Println("d:", valD, "exists?", okD)
	
	if val, ok := points["c"]; ok{
		fmt.Println(val, "found key c")
	}else{
		fmt.Println("key c not found")
	}
	
	prices := map[string]int{
		"apple": 100,
		"banana": 50,
		"orange": 80,
	}
	total := 0
	//looping over maps with for range (to get the key and value)
	for item, price := range prices{
		fmt.Println(item, "price is", price)
		total += price
	}
	
	fmt.Println("Total price:", total)
	
	//looping over maps with for range (to get the key only)
	for item := range prices{
		fmt.Println("Item:", item)
	}
}