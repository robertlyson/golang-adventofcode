package main

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func main() {

}

func distance(p1 Point, p2 Point) float64 {
	return math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y)
}

//https://math.stackexchange.com/a/163101
func spiral(n float64) Point {
	var k = math.Ceil((math.Sqrt(n) - 1) / 2)
	var t = 2*k + 1
	var m = math.Pow(t, 2)
	t = t - 1

	if n >= m-t {
		return Point{k - (m - n), -k}
	}
	m = m - t
	if n >= m-t {
		return Point{-k, -k + (m - n)}
	}
	m = m - t
	if n >= m-t {
		return Point{-k + (m - n), k}
	}

	return Point{k, k - (m - n - t)}
}
