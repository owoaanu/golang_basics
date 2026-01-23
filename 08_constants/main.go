package main

import(
	"fmt"
)


func main(){
	// constat are values that do not change
	const appName ="Go Basics"
	
	//typed constants
	const maxUpload int = 25
	const discountedPrice float64 = 10.3
	
	fmt.Println("App Name:", appName)
	fmt.Println("Max Upload Size (MB):", maxUpload)
	fmt.Println("Discounted Price ($):", discountedPrice)
}