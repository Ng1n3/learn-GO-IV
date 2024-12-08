package main

import (
	"fmt"
	"strconv"
)

type FuncOps struct {
	FirstName string
	LastName  string
	Age       int
}

func myFunc(options FuncOps) {
	fmt.Println(options)
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, val := range a {
		total += int(val)
	}
	return total
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func main() {

	fmt.Println("-----DECLARING AND CALLING FUNCTIONS-----")
	myFunc(FuncOps{FirstName: "Emmanuel", LastName: "Olayinka", Age: 26})

	// ? functions are basically treated as values meaning that you can pass one functino into another function in as much as the function has the same parameter and the same return type.
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result := myFuncVariable("Baby")
	fmt.Println(result)

	myFuncVariable = f2
	result = myFuncVariable("Baby")
	fmt.Println(result)

	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"44", "-", "12"},
		{"33", "+", "11"},
		{"11", "*", "2"},
		{"112", "/", "2"},
		{"22"},
		{"24", "%", "2"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsurpported operator", op)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	// ? ANONYMOUS FUNCTIONS
	f := func(j int) {
		fmt.Println("printing", j, "From inside an ANONYMOUS function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}

  // [lecture] -> CLOSURE
  a := 20
  f3 := func() {
    fmt.Println(a)
    a = 30
    fmt.Println(a)
  }
  f3()
  fmt.Println(a)
}
