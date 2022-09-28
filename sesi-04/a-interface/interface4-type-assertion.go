package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Luas Lingkaran
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Luas Persegi Panjang
func (r rectangle) area() float64 {
	return r.height * r.width
}

// Keliling Lingkaran
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Keliling Persegi Panjang
func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.width)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}

	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}
}
