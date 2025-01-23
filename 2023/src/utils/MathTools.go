package utils

import "math"

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(Gcd(a, b)))
}
