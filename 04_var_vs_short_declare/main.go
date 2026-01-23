package main

import(
	"fmt"
)

func main(){
	var city string
	city = "Auchi"
	
	// leveraging go var inferencing
	var channel = "owoaanu" //inferred to string
	
	// the short way to declare and assign a variable
	// :=  this declares and assigns a variable with type inference
	subscribers := 500
	
	subscribers += 1000
	
	// short way to declare and assign multiple variables
	likes, comments := 100, 30
	
	fmt.Println("City:", city)
	fmt.Println("Channel:", channel)
	fmt.Println("Subscribers:", subscribers)
	fmt.Println("Likes:", likes)
	fmt.Println("Comments:", comments)
}