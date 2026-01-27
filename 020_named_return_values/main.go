package main

import "fmt"

// named return value
func subtractAndDivide(a int, b int) (difference int, quotient int){
	difference = a - b
	quotient = a/b
	
	return
}

func main(){
	difference, quotient := subtractAndDivide(5,15)
	
	fmt.Println("difference:", difference, "quotient", quotient )
}