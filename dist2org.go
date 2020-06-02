package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) dist2Org() float64 {
	res := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	res = math.Sqrt(res)
	return res
}

func main() {
	p := Point{x: 3, y: 4}
	res := p.dist2Org()
	fmt.Println(res)
}
