package closestpoints

import (
	// "fmt"
	"fmt"
	"math"
)

func D_C_util(points Points) (Point, Point) {
	sortedByX := SortByX(points)
	sortedByY := SortByY(points)
	sortedByX.Print()
	fmt.Println("\nHalfway", sortedByX.Points[sortedByX.Len()/2-1])
	return GetClosestPointPair(sortedByX, sortedByY)
}

func GetClosestPointPair(sortedByX, sortedByY Points) (Point, Point) {
	l := sortedByX.Len()
	// fmt.Println("length", l)
	if l == 2 {
		return sortedByX.Points[0], sortedByX.Points[1]
	}

	if l == 3 {
		d_0_1 := getDistance(sortedByX.Points[0], sortedByX.Points[1])
		d_1_2 := getDistance(sortedByX.Points[1], sortedByX.Points[2])
		d_0_2 := getDistance(sortedByX.Points[0], sortedByX.Points[2])

		switch min := math.Min(math.Min(d_0_1, d_1_2), d_0_2); min {
		case d_0_1:
			return sortedByX.Points[0], sortedByX.Points[1]
		case d_1_2:
			return sortedByX.Points[1], sortedByX.Points[2]
		case d_0_2:
			return sortedByX.Points[0], sortedByX.Points[2]
		}

	}

	half := l / 2
	var left_y, right_y Points
	for _, x := range sortedByY.Points {
		if x.X < sortedByX.Points[half-1].X {
			left_y.Points = append(left_y.Points, x)
		} else {
			right_y.Points = append(right_y.Points, x)
		}

	}
	left1, left2 := GetClosestPointPair(Points{Points: sortedByX.Points[:half]}, left_y)
	right1, right2 := GetClosestPointPair(Points{Points: sortedByX.Points[half:]}, right_y)

	d_left := getDistance(left1, left2)
	d_right := getDistance(right1, right2)

	d := math.Min(d_left, d_right)

	var sliced Points
	for _, x := range sortedByY.Points {
		if sortedByX.Points[half-1].X-int(math.Ceil(d)) < x.X && x.X < sortedByX.Points[half-1].X+int(math.Floor(d)) {
			sliced.Points = append(sliced.Points, x)
		}
	}

	if sliced.Len() > 1 {
		// fmt.Println("Testing split cases now")
		// sliced.Print()

		var split_1, split_2 Point
		d_split := d
		for x := 0; x < sliced.Len()-1; x++ {
			for y := 1; y < int(math.Min(7, float64(sliced.Len()-x))); y++ {
				d_new := getDistance(sliced.Points[x], sliced.Points[x+y])
				if d_new < d_split {
					// fmt.Println(d_new, d_split)
					d_split = d_new
					split_1 = sliced.Points[x]
					split_2 = sliced.Points[x+y]
				}
			}
		}
		// fmt.Println("Testing split cases over")
		// fmt.Println("d_left, d_right, d_split")
		// fmt.Println(d_left, d_right, d_split)
		switch m := math.Min(math.Min(d_left, d_right), d_split); m {
		case d_left:
			return left1, left2
		case d_right:
			return right1, right2
		case d_split:
			return split_1, split_2
		}
	}
	if d_left <= d_right {
		return left1, left2
	} else {
		return right1, right2
	}
}
