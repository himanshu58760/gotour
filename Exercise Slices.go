package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) (a [][]uint8) {
	a = make([][]uint8, dy)

	for i := range a {
		a[i] = make([]uint8, dx)
	}

	for i := range a {
		for j := range a[i] {
			switch {
			case j % 7 == 0:
				a[i][j] = 255
			case j % 3 == 0:
				a[i][j] = 128
			case j % 5 == 0:
				a[i][j] = 64
			default:
				a[i][j] = 8
			}
		}
	}

	return
}

func main() {
	pic.Show(Pic)
}