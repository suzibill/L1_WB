package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func newPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func main() {
	p1 := newPoint(1, 2)
	p2 := newPoint(3, 4)
	fmt.Println(p1.Distance(p2))
}
