package main

import (
	"fmt"
)

// creating a type for a function type
type x func (a int, b int) int

// has same function signature as x
func sum (a int, b int) int {
	return a + b
}

// function which accepts another function as argument
func execFunction (f x){

 fmt.Println("val : ", f(1,2))
}

func main() {
	//passing function here
	execFunction(sum)
}
