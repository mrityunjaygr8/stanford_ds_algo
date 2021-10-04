package closestpoints

import (
	"fmt"
	"math"
)

func GetClosestPointPair(points Points) (Point, Point) {
	l := points.Len()
	fmt.Println("length", l)
	if l == 2 {
		return points.Points[0], points.Points[1]
	}

	if l == 3 {
		d_0_1 := getDistance(points.Points[0], points.Points[1])
		d_1_2 := getDistance(points.Points[1], points.Points[2])
		d_0_2 := getDistance(points.Points[0], points.Points[2])

		switch min := math.Min(math.Min(d_0_1, d_1_2), d_0_2); min {
		case d_0_1:
			return points.Points[0], points.Points[1]
		case d_1_2:
			return points.Points[1], points.Points[2]
		case d_0_2:
			return points.Points[0], points.Points[2]
		}

	}

	// p1, p2 := Point{}, Point{}
	fmt.Println("sortedByX")
	sortedByX := SortByX(points)
	sortedByX.Print()

	fmt.Println("sortedByY")
	_ = SortByY(points)
	// sortedByY.Print()

	half := l / 2
	left1, left2 := GetClosestPointPair(Points{Points: sortedByX.Points[:half]})
	right1, right2 := GetClosestPointPair(Points{Points: sortedByX.Points[half:]})

	d_left := getDistance(left1, left2)
	d_right := getDistance(right1, right2)

	d := math.Min(d_left, d_right)

	var sliced Points
	for _, x := range sortedByX.Points {
		if sortedByX.Points[half].X-int(math.Ceil(d)) < x.X && x.X < sortedByX.Points[half].X+int(math.Floor(d)) {
			sliced.Points = append(sliced.Points, x)
		}
	}

	// fmt.Println("Testing split cases now")
	// sliced.Print()

	if sliced.Len() > 1 {

		split_1, split_2 := sliced.Points[0], sliced.Points[sliced.Len()-1]
		d_split := getDistance(split_1, split_2)
		for x := 0; x < sliced.Len()-1; x++ {
			for y := x + 1; y < int(math.Min(7, float64(sliced.Len()-1))); y++ {
				d_new := getDistance(sliced.Points[x], sliced.Points[y])
				if d_new < d_split {
					d_split = d_new
					split_1 = sliced.Points[x]
					split_2 = sliced.Points[y]
				}
			}
		}
		switch m := math.Min(math.Min(d_left, d_right), d_split); m {
		case d_left:
			return left1, left2
		case d_right:
			return right1, right2
		case d_split:
			return split_1, split_2
		}
	}
	// fmt.Println("Testing split cases over")
	if d_left <= d_right {
		return left1, left2
	} else {
		return right1, right2
	}
}
