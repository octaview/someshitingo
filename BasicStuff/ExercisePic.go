package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slave := make([][]uint8, dy)
	for y := 0; y <= dy; y++ {
		slave[y] = make([]uint8, dx)
		for x := 0; x <= dx; x++ {
			slave[y][x] = uint8((x + y) / 2)
		}
	}
	return slave
}

func main() {
	pic.Show(Pic)
}
