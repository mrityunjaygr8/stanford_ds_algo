package multiplication

import (
	"math"
	"strconv"
)

func Karatsuba(num_1, num_2 int) int {
	// fmt.Println(num_1_slice, num_2_slice)
	num := 0
	if num_1 > num_2 {
		num = do_karat(num_1, num_2)
	} else {
		num = do_karat(num_2, num_1)
	}
	return num
}

func do_karat(num_1, num_2 int) int {
	// fmt.Printf("num_1\tnum_2\n")
	// fmt.Printf("%d\t%d\n", num_1, num_2)

	// num_1_slice := bigToDigits(num_1)
	// num_2_slice := bigToDigits(num_2)

	// num_2_len := len(num_2_slice)
	// if num_2_len > 1 {
	// 	m := num_2_len / 2

	// 	a := num_1_slice[:len(num_1_slice)-m]
	// 	b := num_1_slice[len(num_1_slice)-m:]
	// 	c := num_2_slice[:len(num_2_slice)-m]
	// 	d := num_2_slice[len(num_2_slice)-m:]

	// 	a_num := digitsToBig(a)
	// 	b_num := digitsToBig(b)
	// 	c_num := digitsToBig(c)
	// 	d_num := digitsToBig(d)

	// 	ac := Karatsuba(a_num, c_num)
	// 	bd := Karatsuba(b_num, d_num)
	// 	abcd := Karatsuba(a_num+b_num, c_num+d_num)

	// 	result := math.Pow10(2*m)*float64(ac) + math.Pow10(m)*float64(abcd-ac-bd) + float64(bd)
	// 	return int(result)

	// } else {
	// 	return num_1 * num_2
	// }

	if num_1 < 10 && num_2 < 10 {
		return num_1 * num_2
	}

	n := math.Max(float64(len(strconv.Itoa(num_1))), float64(len(strconv.Itoa(num_2))))
	m := math.Ceil(n / 2)

	x_H := math.Floor(float64(num_1) / math.Pow10(int(m)))
	x_L := num_1 % int(math.Pow10(int(m)))

	y_H := math.Floor(float64(num_2) / math.Pow10(int(m)))
	y_L := num_2 % int(math.Pow10(int(m)))

	a := do_karat(int(x_H), int(y_H))
	b := do_karat(x_L, y_L)
	c := do_karat(int(x_H)+x_L, int(y_H)+y_L) - a - b

	res := a*int(math.Pow10(int(m*2))) + c*int(math.Pow10(int(m))) + b
	return res

}
