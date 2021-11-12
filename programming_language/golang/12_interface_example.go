package main

import (
	"fmt"
	"math"
)

// Create 2 struct
type rectangle struct {
	length, width float64
}

type circle struct {
	radius float64
}

// Code an interface
type geometry interface {
	area() float64
	perimeter() float64
}

// Implement the interface for each struct
func (rect *rectangle) area() float64 {
	return rect.length * rect.width
}

func (rect *rectangle) perimeter() float64 {
	return 2*rect.length + 2*rect.width
}

func (cir *circle) area() float64 {
	return math.Pi * cir.radius * cir.radius
}

func (cir *circle) perimeter() float64 {
	return 2 * math.Pi * cir.radius
}

func measure(g geometry) {
	fmt.Printf("[+] Shape: %v, Area %v, Perimeter %v\n", g, g.area(), g.perimeter())
}

func InterfaceExample() {
	fmt.Println("===== Start Interface Example =====")
	rectObject := rectangle{10, 15}
	cirObject := circle{7}

	measure(&rectObject)
	measure(&cirObject)
}
