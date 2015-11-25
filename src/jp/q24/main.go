package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	n := 0
	return loop(x, z, n)
}

func loop(x, z float64, n int) float64 {
	if n >= 10 {
		return z
	} else {
		n += 1
		return loop(x, z-(math.Pow(z, 2)-x)/(2*z), n)
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
