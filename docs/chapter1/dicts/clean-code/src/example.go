package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

var points = []Point{{4, 5}, {7, 1}, {2, 9}}

const maxFloat64 = 1.79769313486231570814527423731704356798070e+308

func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	minDistance := maxFloat64
	origin := Point{0, 0}
	// 入力した点のうち最も原点に近い点を探し、その距離を求める
	for _, p := range points {
		if minDistance < distance(origin, p) {
			minDistance = distance(origin, p)
		}
	}
	fmt.Println(minDistance)
}
