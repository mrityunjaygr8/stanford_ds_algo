package multiplication

import (
	"math"
)

func Basic_multiplication(num_1, num_2 int) int {

	// reverse_start := time.Now()
	num_1_slice := reverse_slice(bigToDigits(num_1))
	num_2_slice := reverse_slice(bigToDigits(num_2))
	// fmt.Println("Basic Reverse Time: ", time.Since(reverse_start))

	sum := 0
	for id_x, x := range num_2_slice {
		tmp_sums := 0
		for id_y, y := range num_1_slice {
			tmp_sums += int(math.Pow(float64(10), float64(id_y))) * (x * y)
		}
		sum += int(math.Pow(float64(10), float64(id_x))) * tmp_sums
	}

	return sum
}
