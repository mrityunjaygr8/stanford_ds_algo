package multiplication

import (
	"math"
)

// "fmt"
// "math"

func bigToDigits(num int) []int {
	num_slice := []int{}
	for num > 0 {
		quo := num / 10
		rem := num % 10

		num_slice = append(num_slice, rem)

		num = quo
	}
	return_slice := reverse_slice(num_slice)
	return return_slice
}

func digitsToBig(nums []int) int {
	l := len(nums)
	cop := make([]int, len(nums))
	_ = copy(cop, nums)

	num := reverse_slice(cop)
	big := 0

	for x := 0; x < l; x++ {
		big += num[x] * int(math.Pow10(x))
	}
	return big
}

func reverse_slice(num_slice []int) []int {
	for i, j := 0, len(num_slice)-1; i < j; i, j = i+1, j-1 {
		num_slice[i], num_slice[j] = num_slice[j], num_slice[i]
	}
	return num_slice
}
