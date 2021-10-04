package matric

import (
	"fmt"
	"math/rand"
	"os"
)

type Matrix [][]int

func (m1 Matrix) add(m2 Matrix) Matrix {
	var res Matrix
	for x := range m1 {
		var col []int
		for y := range m2 {
			col = append(col, m1[x][y]+m2[x][y])
		}
		res = append(res, col)
	}
	return res
}

func (m1 Matrix) sub(m2 Matrix) Matrix {
	var res Matrix
	for x := range m1 {
		var col []int
		for y := range m2 {
			col = append(col, m1[x][y]-m2[x][y])
		}
		res = append(res, col)
	}
	return res
}

func Breakup(m1 Matrix) (Matrix, Matrix, Matrix, Matrix) {
	var a, b, c, d Matrix
	length := len(m1)
	for _, x := range m1 {
		if length != len(x) {
			fmt.Println("This is not a square array")
			os.Exit(-1)
		}
	}

	h := length / 2

	for _, x := range m1[:h] {
		var cols []int
		for _, y := range x[:h] {
			cols = append(cols, y)
		}
		a = append(a, cols)
	}
	for _, x := range m1[h:] {
		var cols []int
		for _, y := range x[h:] {
			cols = append(cols, y)
		}
		d = append(d, cols)
	}
	for _, x := range m1[:h] {
		var cols []int
		for _, y := range x[h:] {
			cols = append(cols, y)
		}
		b = append(b, cols)
	}
	for _, x := range m1[h:] {
		var cols []int
		for _, y := range x[:h] {
			cols = append(cols, y)
		}
		c = append(c, cols)
	}

	return a, b, c, d
}

func backTogether(a, b, c, d Matrix) Matrix {
	var m Matrix

	for x := range a {
		m = append(m, append(a[x], b[x]...))
	}

	for x := range c {
		m = append(m, append(c[x], d[x]...))
	}

	return m
}

func PrintMatrix(m Matrix) {
	for _, x := range m {
		fmt.Println(x)
	}
}

func CreateArrays(size int) (Matrix, Matrix) {
	const RAND = 1024
	var A Matrix
	var B Matrix
	for i := 0; i < size; i++ {
		col1 := []int{}
		col2 := []int{}
		for j := 0; j < size; j++ {
			col1 = append(col1, rand.Intn(RAND))
			col2 = append(col2, rand.Intn(RAND))
		}
		A = append(A, col1)
		B = append(B, col2)
	}
	return A, B
}
