package closestpoints

// import "fmt"

func Naive(points Points) (Point, Point) {
	l := points.Len()
	p1, p2 := points.Points[0], points.Points[l-1]
	d := getDistance(points.Points[0], points.Points[l-1])

	for x := 0; x < l-1; x++ {
		for y := x + 1; y < l; y++ {
			d_new := getDistance(points.Points[x], points.Points[y])
			// fmt.Println(points.Points[x], points.Points[y], d_new)
			if d_new < d {
				p1 = points.Points[x]
				p2 = points.Points[y]
				d = d_new
			}
		}
	}
	return p1, p2
}
