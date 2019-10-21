package student

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for y1 := 0; y1 < y; y1++ {
		for x1 := 0; x1 < x; x1++ {
			if (y1 == 0 && x1 == 0) || (y != 1 && y1 == y-1 && x1 == x-1 && x != 1) {
				z01.PrintRune('A')
			} else if (y1 == 0 && x1 == x-1) || (y1 == y-1 && x1 == 0) {
				z01.PrintRune('C')
			} else if x1 == x-1 || x1 == 0 {
				z01.PrintRune('B')
			} else if y1 == y-1 || y1 == 0 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(10)
	}
}
