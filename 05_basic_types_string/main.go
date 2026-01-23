package main

import(
	"fmt"
	"strings"
)

func main(){
	firstName := "Samuel"
	lastName := "Adebayo-Adesina"
	fullName := firstName + " " + lastName
	
	fmt.Println("Full Name:", fullName)
	
	fmt.Println(strings.ToUpper(fullName))
}