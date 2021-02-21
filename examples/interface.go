package main

import (
	"fmt"
	"math"
)

const pi = 3.14151692

// The interface
type geometry interface {
	area() float64
	perim() float64
}

// rect will implement the methods of the geometry interface
type rect struct {
	width, height float64
}

// circle will implement the methods of geometry interface
type circle struct {
	radius float64
}

// rect area implementation
func (r rect) area() float64 {
	return r.height * r.width
}

// rect perim implementation
func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

// circle area implementation
func (c circle) area() float64 {
	return pi * math.Pow(c.radius, 2)
}

// circle perim implementation
func (c circle) perim() float64 {
	return 2 * pi * c.radius
}

func mesure(g geometry) {
	fmt.Printf("%T\n", g)
	fmt.Printf("\tarea:%v\n", g.area())
	fmt.Printf("\tperim:%v\n", g.perim())
}

func main() {
	r := rect{
		width:  3,
		height: 4,
	}
	mesure(r)

	c := circle{
		radius: 5,
	}
	mesure(c)
}
