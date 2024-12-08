package main

import (
	"fmt"
	"maps"
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

	// [lecture] -> This means elements of the original slice beyond the end of the subslice, including unused capacity, are shared by both slices. You can also slice arrays too.

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
	fmt.Println()

	fmt.Println("---converting arrays to slices----")
	// ? -> Taking a slice from an array has the same memory-sharing properties as taking a slice from a slice.
	c := [4]int{5, 6, 7, 8}
	d := c[:2]
	e := c[2:]

	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println()

	fmt.Println("-----Converting Slices to Arrays-----")
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 14
	fmt.Println("xSLice: ", xSlice)
	fmt.Println("xArray: ", xArray)
	fmt.Println("smallArray: ", smallArray)
	fmt.Println()

	fmt.Println("-----Strings and Runes and Bytes------")
	var s1 string = "Hello 🌑"
	var s2 string = s1[4:7]
	var s3 string = s1[:5]
	var s4 string = s1[6:]
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)
	fmt.Println(len(s1))
	fmt.Println()

	fmt.Println("-----MAPS-----")
	f := map[string]int{
		"hello": 9,
		"world": 15,
		"!":     10,
		"...":   20,
	}

	g, ok := f["hello"]
	fmt.Println(g, ok)
	h, ok := f["baddie"]
	fmt.Println(h, ok)

	// ? DELETING FROM MAPS
	delete(f, "hello")
	i, ok := f["hello"]
	fmt.Println(i, ok)

	// ? EMPTYING MAPS
	clear(f)
	fmt.Println(f, len(f))

	// ? COMPARING MAPS
	j1 := map[string]int{
		"hello": 9,
		"world": 15,
		"!":     10,
		"...":   20,
	}

	j2 := map[string]int{
		"hello": 9,
		"world": 15,
		"!":     10,
		"...":   20,
	}
	fmt.Println(maps.Equal(j1, j2))

	// ? USING MAPS AS SETS
	intSet := map[int]bool{}
	vals := []int{5, 10, 5, 20, 5, 45, 100, 1, 2}
	for _, v1 := range vals {
		intSet[v1] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	fmt.Println()

	fmt.Println("-----STRUCTS-----")
	type person struct {
		name string
		age  int
		car  string
	}

	var emmanuel person

	emmanuel = person{
		"emmanuel",
		26,
		"Camry",
	}

	fmt.Println(emmanuel)

	// ? ANONYMOUS STRUCTS
	car := struct {
		name string
		year int
	}{
		name: "Toyota Camry",
		year: 2016,
	}

	fmt.Println(car.name, car.year)

	// ! -> Go doesn't allow comparison between variables that represent structs of different types

	type firstPerson struct {
		name string
		age  int
	}

	f1 := firstPerson{
		name: "Emmanuel",
		age:  26,
	}

	type secondPerson struct {
		name string
		age  int
	}

	f2 := secondPerson{
		name: "Daniella",
		age:  26,
	}

	//? --> you can use a type conversionto convert te instance of firstPerson to secondPerson, but you cannot use == to compare the instance of firstPerson and an instance of secondPerson, because they are of different types.
	// fmt.Println("comparing f1 and f2", f1 == f2)
	f2 = secondPerson(f1)
	fmt.Println(f2, f1);
	

	type thirdPerson struct {
		age  int
		name string
	}

	f3 := thirdPerson {
		age: 26,
		name: "Emmanuel",
	}

	// ? --> you can't convert an instance of first person to thirdPerson, because the field are in different order

	type fourthPerson struct {
		firstName string
		age       int
	}

	f4 := fourthPerson {
		firstName: "Emmanuel",
		age: 26,
	}

	// ? --> you cannot convert an instance of firstPerson to fourthPerson because the names don't match

	type fifthPerson struct {
		name          string
		age           int
		favoritecolor string
	}
	f5 := fifthPerson {
		name: "Emmanuel",
		age: 26,
		favoritecolor: "Red",
	}

	// ? --> you cannot convert an instance of firstPerson to fifthPerson because there is an additional field

}
