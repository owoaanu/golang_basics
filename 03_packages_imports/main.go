package main

// making imports, brings external packages into the current file
import(
	"fmt"
	"math"
)

func main(){
	// to use an external package
	// packageName.FunctionName -> call a function from the package
	
	fmt.Println("sqrt(25)", math.Sqrt(25)) 
}