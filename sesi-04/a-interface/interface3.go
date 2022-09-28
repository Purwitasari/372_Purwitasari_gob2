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

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
	//mengikuti (*)
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	//(*)
	print("Rectangle", c1)
	print("Circle", r1)

}
