package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	x := make([]uint8, dx)
	y := make([][]uint8, dx, dy)

	for i := range y {
		x[i] = uint8(i)
		for j := range y[i] {
			y[i][j] = uint8((i + j) / 2)
		}
	}
	return y

}

func main() {
	pic.Show(Pic)
}
