package matric

import (
	"fmt"
	"os"
)

func Strassen(m1, m2 Matrix) Matrix {
	// fmt.Println("being strassen")
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

	p1 := Strassen(a.add(d), e.add(h))
	p2 := Strassen(c.add(d), e)
	p3 := Strassen(a, f.sub(h))
	p4 := Strassen(d, g.sub(e))
	p5 := Strassen(a.add(b), h)
	p6 := Strassen(c.sub(a), e.add(f))
	p7 := Strassen(b.sub(d), g.add(h))

	A := p1.add(p4).sub(p5).add(p7)
	B := p3.add(p5)
	C := p2.add(p4)
	D := p1.sub(p2).add(p3).add(p6)

	m := backTogether(A, B, C, D)
	return m
}
