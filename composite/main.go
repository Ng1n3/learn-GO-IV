package main

import (
	"fmt"
	"slices"
)

func main() {
	var x = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(x)

	var j = []int{1, 2, 3, 4}
	var k = []int{1, 2, 3, 4}
	var l = []string{"a", "red", "ball"}
	var m = []string{"a", "red", "ball"}

	fmt.Println(slices.Equal(j, k))
	fmt.Println(slices.Equal(l, m))
	// fmt.Println(slices.Equal(j, m))

	j = append(j, k...) //? appending two slices
	j = append(j, 6, 7, 8, 9)

	fmt.Println(j)

	// ? -> Mkae is used to create a slice with inital capaicty

	n := make([]int, 5) // ? -> ceeates an int slice with length of 5 and capacity of 5
	n = append(n, 20)   // ? -> append always increases the size of the slice not its 6

	fmt.Println(n)

	clear(n)

	fmt.Println(n, len(n))

	o := []string{"a", "b", "c", "d", "e"}
	p := o[:2]
	q := o[:1]
	r := o[1:3]
	s := o[:]
	fmt.Println("o: ", o)
	fmt.Println("p: ", p)
	fmt.Println("q: ", q)
	fmt.Println("r: ", r)
	fmt.Println("s: ", s)

	//! -> changes made to an element in a slice affects all slices that share that element

	t := []string{"a", "b", "c", "d", "e"}
	u := t[:2]
	v := t[1:]

	t[1] = "u"
	u[0] = "t"
	v[1] = "v"

	fmt.Println("t: ", t)
	fmt.Println("u: ", u)
	fmt.Println("v: ", v)

	//! -> Avoid combining append wth slicing slices

	// [lecture] -> This means elements of the original slice beyond the end of the subslice, including unused capacity, are shared by both slices.

	w := make([]string, 0, 5)
	w = append(w, "a", "b", "c", "d")
	y := w[:2:2]
	z := x[2:4:4]

	fmt.Println("---------------")
	fmt.Println()
	fmt.Println("w: ", w)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println()


  fmt.Println("------COPY------")
  //[lecture] -> COPY used to create a slice that is independent of the original. copy(destination slice, source slice)

  a := []int{1, 2, 3, 4}
  b := make([]int, 4)
  num := copy(b, a)
  fmt.Println(b, num)
}
