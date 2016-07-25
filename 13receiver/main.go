package main

import (
	"fmt"
	"math"
)

// Point represents a point in space.
type Point struct {
	X int
	Y int
}

// New returns a Point based on X and Y positions on a graph.
func New(x int, y int) Point {
	return Point{x, y}
}

// Distance finds the length of the hypotenuse between two points.
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}

func main() {
	p1 := New(37, -76)
	p2 := New(26, -80)

	dist := p1.Distance(p2)
	fmt.Println("Distance", dist)
}
