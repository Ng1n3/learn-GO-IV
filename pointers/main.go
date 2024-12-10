package main

import "fmt"

func main() {
	// [lecture] -> A pointer is a variable that contains the address where another variable is stored.

	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

  //? -> The only time you should use pointer parameters to modify a variable is when the function expects an interface.For data structures that are smaller than 10 megabytes, it is actually slower to return a pointer type than a value type.


  
}
