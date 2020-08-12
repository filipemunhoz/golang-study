package main

import "fmt"

const (
	_     = iota // first value ignored
	a     = iota
	b     = iota
	c int = iota
	d     = iota + 10 // untyped int
	e     = iota + 20
	f     = iota + 30
	g     = iota + 4.4 // untyped float
	h     = iota + 5.5
)

const (
	i float64 = iota // reseted iota
)

// Size int - using enum type
type Size int

// Const for enum
const (
	Small Size = iota
	Medium
	Big
	Huge
)

func main() {

	fmt.Printf("%v - %T\n", a, a)
	fmt.Printf("%v - %T\n", b, b)
	fmt.Printf("%v - %T\n", c, c)
	fmt.Printf("%v - %T\n", d, d)
	fmt.Printf("%v - %T\n", e, e)
	fmt.Printf("%v - %T\n", f, f)
	fmt.Printf("%v - %T\n", g, g)
	fmt.Printf("%v - %T\n", h, h)
	fmt.Printf("%v - %T\n", i, i)

	var s Size = Medium
	fmt.Println("Size: ", s)
}

func (d Size) String() string {
	return [...]string{"Small", "Medium", "Big", "Huge"}[d]
}
