package student

import "github.com/01-edu/z01"

func Raid1a(x,y int) {	
	for y1 := y; y1 > 0; y1-- {
		for x1 := x; x1 > 0; x1-- {
			if (y1 == y && x1 == x) || (y1 == 1 && x1 == 1) || (y1 == y && x1 == 1) || (y1 == 1 && x1 == x)  {
				z01.PrintRune('o')
			} else if x1 == x || x1 == 1 {
				z01.PrintRune('|')
			} else if y1 == y || y1 == 1 {
				z01.PrintRune('-')
			} else {
				z01.PrintRune(' ')
	}
			}
	}
		}
	}
		z01.PrintRune(10)
	}
}
