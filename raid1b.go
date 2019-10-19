package student

import (
	"github.com/01-edu/z01"
)

func startEnd(a, b int) {
	z01.PrintRune('/')
	for i := a - 3; i >= 0; i-- {
		z01.PrintRune('*')
	}
	z01.PrintRune(92)
	z01.PrintRune(10)
}

func startEnd0(a, b int) {
	z01.PrintRune(92)
	for i := a - 3; i >= 0; i-- {
		z01.PrintRune('*')
	}
	z01.PrintRune('/')
	z01.PrintRune(10)
}

func middle(a, b int) {
	if b > 2 {
		z01.PrintRune('*')
		for i := a - 3; i >= 0; i-- {
			z01.PrintRune(' ')
		}
		z01.PrintRune('*')
		z01.PrintRune(10)
	}
}

func middle1(a, b int) {
	if b == 1 {
		z01.PrintRune('/')
		z01.PrintRune(10)
	}
	if b == 2 {
		z01.PrintRune('/')
		z01.PrintRune(10)
		z01.PrintRune(92)
		z01.PrintRune(10)
	}
	if b > 2 {
		z01.PrintRune('/')
		z01.PrintRune(10)
		for i := b - 2; i > 0; i-- {
			z01.PrintRune('*')
			z01.PrintRune(10)
		}
		z01.PrintRune(92)
		z01.PrintRune(10)
	}
}

func Raid1b(a, b int) {
	if a >= 2 && b > 1 {
		startEnd(a, b)
		for i := b - 2; i > 0; i-- {
			middle(a, b)
		}
		startEnd0(a, b)
	} else if a >= 2 && b == 1 {
		startEnd(a, b)
	} else if a == 1 {
		middle1(a, b)
	} else if a < 0 || b < 0 {
		return 
	}
}
