// Exercise Link: https://go.dev/tour/flowcontrol/8
// Description: Compute the square root of a number using Newton's method

package main

import (
	"fmt"
	"math"
)

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(z)
// 		z -= (z * z - x) / (2 * z)
// 	}
// 	return z
// }

func Sqrt(x float64) float64 {
	var z0, zx, diff = 1.0, 1.0, 100.0

	for i := 0; diff > 0.00001; i++ {
		fmt.Println(zx)
		zx = z0 - (z0 * z0 - x) / (2 * z0)
		diff = math.Abs(zx - z0)
		z0 = zx
	}
	return z0
}

func main() {
	fmt.Println(Sqrt(2))
}
