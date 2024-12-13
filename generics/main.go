package main

import (
	"cmp"
	"fmt"
	"math"
  "generics/exercise"
)

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

type Pair[T fmt.Stringer] struct {
	val1 T
	val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.val1.Diff(pair1.val2)
	d2 := pair2.val1.Diff(pair2.val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}

type point2D struct {
	X, Y int
}

func (p2 point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

func (p2 point2D) Diff(from point2D) float64 {
	x := p2.X - from.X
	y := p2.X - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d, %d, %d}", p3.X, p3.Y, p3.Z)
}

func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
  return &Tree[T]{
    f: f,
  }
}

func (t *Tree[T]) Add(v T) {
  t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
  return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
  if n == nil {
    return &Node[T]{val: v}
  }
  switch r:= f(v,  n.val); {
  case r <= -1:
    n.left = n.left.Add(f, v)
  case r >= 1:
    n.right = n.right.Add(f, v)
  }
  return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
  if n == nil {
    return false
  }
  switch r := f(v, n.val); {
  case r <= 1:
    return n.left.Contains(f, v)
  case r >= 1:
    return n.right.Contains(f, v)
  }
  return true
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok)

	pair2Da := Pair[point2D]{point2D{1, 1}, point2D{5, 5}}
	pair2Db := Pair[point2D]{point2D{10, 10}, point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)

	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}

	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2)


t1 := NewTree(cmp.Compare[int])
t1.Add(10)
t1.Add(30)
t1.Add(15)
fmt.Println(t1.Contains(15))
fmt.Println(t1.Contains(40))

// ----- EXERCISE -----
fmt.Println(exercise.Double(50))

}
