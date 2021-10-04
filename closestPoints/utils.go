package closestpoints

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X int
	Y int
}

type Points struct {
	Points []Point
}

func (p Points) Print() {
	for _, x := range p.Points {
		x.Print()
	}
}

func (p Points) Len() int {
	return len(p.Points)
}

func (p Point) Print() {
	fmt.Printf("(%d, %d)\n", p.X, p.Y)
}

func CreatePoints(count int) Points {
	rand.Seed(time.Now().UnixNano())
	points := make([]Point, count)
	for x := 0; x < count; x++ {
		points[x] = Point{X: rand.Intn(255), Y: rand.Intn(255)}
	}

	return Points{Points: points}
}

func SortByX(points Points) Points {
	l := len(points.Points)

	if l <= 1 {
		return points
	}

	half := int(math.Ceil(float64(l) / 2))

	sorted_1 := SortByX(Points{Points: points.Points[:half]})
	sorted_2 := SortByX(Points{Points: points.Points[half:]})

	sorted := []Point{}

	x, y := 0, 0
	for x < len(sorted_1.Points) && y < len(sorted_2.Points) {
		if sorted_1.Points[x].X < sorted_2.Points[y].X {
			sorted = append(sorted, sorted_1.Points[x])
			x++
		} else {
			sorted = append(sorted, sorted_2.Points[y])
			y++
		}
	}

	if x < len(sorted_1.Points) {
		sorted = append(sorted, sorted_1.Points[x:]...)
	}

	if y < len(sorted_2.Points) {
		sorted = append(sorted, sorted_2.Points[y:]...)
	}

	return Points{Points: sorted}
}
func SortByY(points Points) Points {
	l := len(points.Points)

	if l <= 1 {
		return points
	}

	half := int(math.Ceil(float64(l) / 2))

	sorted_1 := SortByY(Points{Points: points.Points[:half]})
	sorted_2 := SortByY(Points{Points: points.Points[half:]})

	sorted := []Point{}

	x, y := 0, 0
	for x < len(sorted_1.Points) && y < len(sorted_2.Points) {
		if sorted_1.Points[x].Y < sorted_2.Points[y].Y {
			sorted = append(sorted, sorted_1.Points[x])
			x++
		} else {
			sorted = append(sorted, sorted_2.Points[y])
			y++
		}
	}

	if x < len(sorted_1.Points) {
		sorted = append(sorted, sorted_1.Points[x:]...)
	}

	if y < len(sorted_2.Points) {
		sorted = append(sorted, sorted_2.Points[y:]...)
	}

	return Points{Points: sorted}
}

func getDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2))
}
