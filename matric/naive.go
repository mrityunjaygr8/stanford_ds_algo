package matric

import (
	"fmt"
	"os"
)

func Naive(m1, m2 Matrix) Matrix {
	// fmt.Println("being naive")
	l1, l2 := len(m1), len(m2)

	if !(l1 == l2) {
		fmt.Println("Matrices are not of equal size")
		os.Exit(-1)
	}

	if l1 == 1 {
		return Matrix{{m1[0][0] * m2[0][0]}}
	}

	a, b, c, d := Breakup(m1)
	e, f, g, h := Breakup(m2)

	p1 := Naive(a, e).add(Naive(b, g))
	p2 := Naive(a, f).add(Naive(b, h))
	p3 := Naive(c, e).add(Naive(d, g))
	p4 := Naive(c, f).add(Naive(d, h))
	m := backTogether(p1, p2, p3, p4)
	return m

}
