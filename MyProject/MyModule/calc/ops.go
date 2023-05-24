package calc

import (
	"fmt"
	"math"
)

func Add(a float64, b float64) float64 {
	return a + b
}
func Minus(a float64, b float64) float64 {
	return a - b
}
func Mul(a float64, b float64) float64 {
	return a * b
}
func Div(a float64, b float64) float64 {
	return a / b
}
func Sqroot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root for negative numbers do not exists")
	}
	return math.Sqrt(a), nil
}
func Pow(a float64, b float64) float64 {
	return math.Pow(a, b)
}
