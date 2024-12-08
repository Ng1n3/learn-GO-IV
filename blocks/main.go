package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// ? --> IF STATEMENTS

	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
		fmt.Println()
	}
	// fmt.Println(n)

	// [LECTURE] -> 4 types of for loops
	/*
	  ? -> complete for loops
	  ? -> The condition only for loops
	  ? -> THe infinite for statement
	  ? -> The for-range statement
	*/

	// ? Complete for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// ? Condition-Only for statement
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// ? Infinite for Statement
	// for {
	// 	fmt.Println("Hello")
	// }

	// ? Using continue to make code clearer
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	//? The for-range statement
	hobbies := []string{"Coding", "video Games", "Tech Tubing", "Business"}
	for i, hobbie := range hobbies {
		fmt.Println(i, hobbie)
	}

	games := map[string]float64{
		"God of War": 9.4,
		"Elden Ring": 9.2,
		"Sekiro":     8.7,
		"EA FC 24":   6.8,
	}

	//? -> Printing the key only.
	for value := range games {
		fmt.Println(value)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range games {
			fmt.Println(k, v)
		}
	}

	// [lecture] --> Iterating over strings
	samples := []string{"Common", "Gee_^!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// ? Whenever a for-range loop encounters a multibyte rune in a string, it converts the UTF-8 representation into a single 32-bit number and assigns it to the value. The offset is incremented by the number of bytes in the rune. If the for-range loop encounters a byte that doesn’t represent a valid UTF-8 value, the Unicode replacement character (hex value 0xfffd) is returned instead.

	// ? LABELLING FOR LOOPS

	cars := []string{"Camry", "Avalon", "Corolla", "Yaris", "Matrix"}
outer:
	for _, car := range cars {
		for i, r := range car {
			fmt.Println(i, r, string(r))
			if r == 'a' {
				continue outer
			}
		}
		fmt.Println()
	}

	for _, car := range cars {
		switch car {
		case "Camry":
			fmt.Println("Dream car")
		case "Corolla":
			fmt.Println("Yeah, not bad at all")
		default:
			fmt.Println("No please!")
		}
	}
}
