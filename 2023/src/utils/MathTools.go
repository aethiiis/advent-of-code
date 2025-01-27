package utils

import (
	"math"
)

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(Gcd(a, b)))
}

func Shoelace(points []Pos) int {
	n := len(points)
	if n < 2 {
		return 0
	}
	area := 0
	for i := 0; i < n-1; i++ {
		area += points[i].R*points[i+1].C - points[i+1].R*points[i].C
	}
	area += points[n-1].R*points[0].C - points[n-1].C*points[0].R
	return int(math.Abs(float64(area / 2)))
}
