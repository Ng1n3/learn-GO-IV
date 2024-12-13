package main

import (
	"fmt"
	"time"
)

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String2() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

const (
	Field1 = 0
  Field2 = 1 + iota
  Field3 = 20
  Field4
  Field5 = iota
)

func main() {

	var c Counter //? c := &Counter{}
	fmt.Println(c.String2())
	c.Increment()
	fmt.Println(c.String2())

	p := Person{
		FirstName: "Emmanuel",
		LastName:  "Olayinka",
		Age:       26,
	}

	output := p.String()
	fmt.Println("output: ", output)

  fmt.Println(Field1, Field2, Field3, Field4, Field5)

}

// ? Go allows you to declare a type at any block level, from the package block down. However, you can access the type only from within its scope. Functions can be defined inside any block. You can use the same method names for different types, but you can't use the same method name for two different methods on the same type.

//? Go uses parameters of pointer type to indicate that a parameter might be modified by the function.

//? When  you use a pointer receiver with a local variable that's a value type, Go automatically takes the address of the local variable when calling the method.

//? Methods in Go are so much like functions that you can use a method as a replacement for a function anytime there's a variable or parameter of a function type.

// [lecture] -> FUNCTION VS Methods
//? -> Anytime your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct, and that logic should be implemented as a method. If your logic depends only on the input parameters, it should be a function

// [lecture] -> iota
//? iota lets you assign an increasing value to a set of constraints
