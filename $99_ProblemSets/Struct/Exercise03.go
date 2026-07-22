package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func inputPoint(p *Point) {
	fmt.Print("Input x and y: ")
	fmt.Scanf("%f %f", &p.x, &p.y)
}

func printPoint(p Point) {
	fmt.Printf("(%.f,%.f)\n", p.x, p.y)
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func closestToRoot(points []Point) Point {
	closestPoint := points[0]
	rootPoint := Point{0, 0}
	closestDistance := distance(rootPoint, points[0])
	for _, point := range points {
		if distance(rootPoint, point) < closestDistance {
			closestDistance = distance(rootPoint, point)
			closestPoint.x = point.x
			closestPoint.y = point.y
		}
	}
	return closestPoint
}

func main() {

	var (
		p1 Point = Point{1, 3}
		p2 Point = Point{3, 4}
		p3 Point = Point{-2, 1}
		p4 Point = Point{2, -2}
	)

	pointList := []Point{p1, p2, p3, p4}
	fmt.Print("Closest point to the root: ")
	printPoint(closestToRoot(pointList))

}
