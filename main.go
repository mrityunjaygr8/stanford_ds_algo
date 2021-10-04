package main

import (
	"fmt"

	"github.com/mrityunjaygr8/stanford_ds_algo/matric"

	"time"
)

func main() {
	// Karatsuba Start
	// num_1 := 12344321
	// num_2 := 43211234
	// basic_start := time.Now()
	// num := m.Basic_multiplication(num_1, num_2)
	// fmt.Println("Basic time taken: ", time.Since(basic_start))
	// fmt.Printf("Basic: %d * %d = %d\n", num_1, num_2, num)
	// karat_start := time.Now()
	// num = m.Karatsuba(num_1, num_2)
	// fmt.Println("Karat time taken: ", time.Since(karat_start))
	// fmt.Printf("Karat: %d * %d = %d\n", num_1, num_2, num)
	// Karatsuba End

	// Inversion Start
	// nums := []int{3, 2, 1, 4, 7, 6, 5, 345, 123, 546, 123, 4456, 2344645, 456}
	// sorted_nums := sorting.MergeSort(nums)
	// fmt.Println("unsorted\t", nums)
	// fmt.Println("sorted\t\t", sorted_nums)

	// nums := []int{1, 3, 5, 8, 2, 4, 6}
	// _, total_inversions := sorting.SplitInversionCount(nums)
	// fmt.Println(total_inversions)
	// Inversion End

	// Strassen's Sub Cubic Matrix Multuplication start
	// s1, s2 := matric.Matrix{{2, 3, 4, 5, 6, 7, 8, 9}, {3, 4, 5, 6, 7, 8, 9, 0}, {4, 5, 6, 7, 8, 9, 0, 1}, {5, 6, 7, 8, 9, 0, 1, 2}, {6, 7, 8, 9, 0, 1, 2, 3}, {7, 8, 9, 0, 1, 2, 3, 4}, {8, 9, 0, 1, 2, 3, 4, 5}, {9, 0, 1, 2, 3, 4, 5, 6}}, matric.Matrix{{3, 4, 5, 6, 7, 8, 9, 0}, {4, 5, 6, 7, 8, 9, 0, 1}, {5, 6, 7, 8, 9, 0, 1, 2}, {6, 7, 8, 9, 0, 1, 2, 3}, {7, 8, 9, 0, 1, 2, 3, 4}, {8, 9, 0, 1, 2, 3, 4, 5}, {9, 0, 1, 2, 3, 4, 5, 6}, {0, 1, 2, 3, 4, 5, 6, 7}}
	// s1, s2 := matric.Matrix{{1, 1}, {1, 1}}, matric.Matrix{{1, 1}, {1, 1}}
	// s1, s2 := matric.Matrix{{2, 3}, {3, 4}}, matric.Matrix{{3, 4}, {4, 5}}
	s1, s2 := matric.CreateArrays(512)
	matric.PrintMatrix(s1)
	matric.PrintMatrix(s2)
	t := time.Now()
	r := matric.Naive(s1, s2)
	fmt.Println("time taken", time.Since(t))
	matric.PrintMatrix(r)
	t1 := time.Now()
	r1 := matric.Strassen(s1, s2)
	fmt.Println("time taken", time.Since(t1))
	matric.PrintMatrix(r1)
	// m2 := matric.Matrix{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}, {4, 5, 6, 7}}
	// matric.PrintMatrix(m2)
	// a, b, c, d := matric.Breakup(m2)
	// fmt.Println("a\t")
	// matric.PrintMatrix(a)
	// fmt.Println("b\t")
	// matric.PrintMatrix(b)
	// fmt.Println("c\t")
	// matric.PrintMatrix(c)
	// fmt.Println("d\t")
	// matric.PrintMatrix(d)
	// fmt.Println(small1, small2, small3, small4)
	// m3 := matric.Matrix{{1, 2, 3, 4, 5}, {2, 3, 4, 5, 6}, {3, 4, 5, 6, 7}, {4, 5, 6, 7, 8}}
	// small3, small4 := matric.Breakup(m3)
	// fmt.Println(small3, small4)

	// Strassen's Sub Cubic Matrix Multuplication end
}
