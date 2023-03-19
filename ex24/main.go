package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

//func (p *Point) Construct(x, y int) {
//	p.x = x
//	p.y = y
//}

func newPoint(x, y int) *Point {
	return &Point{x, y}
}

func main() {
	p1 := newPoint(1, 2)
	p2 := newPoint(3, 4)
	fmt.Println(findDistance(p1, p2))
}

func findDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}
