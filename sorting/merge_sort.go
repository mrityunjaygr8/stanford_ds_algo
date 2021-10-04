package sorting

import (
	"fmt"
	"math"
)

func MergeSort(nums []int) []int {
	l := len(nums)

	if l <= 1 {
		return nums
	}

	half := int(math.Ceil(float64(l) / 2))

	sorted_1 := MergeSort(nums[:half])
	sorted_2 := MergeSort(nums[half:])

	sorted := []int{}

	x, y := 0, 0
	for x < len(sorted_1) && y < len(sorted_2) {
		if sorted_1[x] < sorted_2[y] {
			sorted = append(sorted, sorted_1[x])
			x++
		} else {
			sorted = append(sorted, sorted_2[y])
			y++
		}
	}

	if x < len(sorted_1) {
		sorted = append(sorted, sorted_1[x:]...)
	}

	if y < len(sorted_2) {
		sorted = append(sorted, sorted_2[y:]...)
	}

	return sorted
}

func SplitInversionCount(nums []int) ([]int, int) {
	l := len(nums)

	if l <= 1 {
		return nums, 0
	}

	half := int(math.Ceil(float64(l) / 2))

	sorted_1, inversions_1 := SplitInversionCount(nums[:half])
	sorted_2, inversions_2 := SplitInversionCount(nums[half:])
	fmt.Println("sorted_1", sorted_1, inversions_1)
	fmt.Println("sorted_2", sorted_2, inversions_2)

	total_inversions := inversions_1 + inversions_2

	sorted := []int{}

	x, y := 0, 0
	fmt.Println(total_inversions)
	for x < len(sorted_1) && y < len(sorted_2) {
		if sorted_1[x] < sorted_2[y] {
			sorted = append(sorted, sorted_1[x])
			x++
		} else {
			sorted = append(sorted, sorted_2[y])
			total_inversions += len(sorted_1) - x
			y++
		}
	}

	if x < len(sorted_1) {
		sorted = append(sorted, sorted_1[x:]...)
	}

	if y < len(sorted_2) {
		sorted = append(sorted, sorted_2[y:]...)
	}

	return sorted, total_inversions

}
