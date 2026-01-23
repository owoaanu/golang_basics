package main

import (
	"fmt"
)

func main() {
	isLogged := true // inferred as boolean
	isAdmin := false
	hasSubscription := true

	// AND &&
	canOpenDashboard := isLogged && hasSubscription

	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println("Can Open Dashboard:", canOpenDashboard)
	fmt.Println("Can Delete Post:", canDeletePost)

	age := 25
	isAdult := age > 18
	
	fmt.Println("User is Adult:", isAdult)

}
