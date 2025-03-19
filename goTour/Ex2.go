// Exercise Link: https://go.dev/tour/moretypes/18
// Description:
// implement a function that returns a slice of length dy, each element of which 
// is a slice of dx 8-bit unsigned integers. When you run the program, 
// it will display your picture, interpreting the integers as grayscale 
// (with the value 0 as black and the value 255 as white).


package main

import "golang.org/x/tour/pic"

func pixel(x, y int) uint8 {
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i, _ := range pic {
		pic[i] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			pic[y][x] = pixel(x, y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
