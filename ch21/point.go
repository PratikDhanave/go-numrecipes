package ch21

import (
	"math"
)

const (
	N = 2
)

type Point [N]float64

// Dist returns the distance between 2 points.
func Dist(p1, p2 Point) float64 {
	dist := 0.0
	for i := 0; i < N; i++ {
		x := p1[i] - p2[i]
		dist += x * x
	}
	return math.Sqrt(dist)
}

// Box represents a rectangular region of 2D space.
type Box struct {
	Min Point
	Max Point
}

func (box Box) DistToPoint(p Point) float64 {
	dist := 0.0

	for i := 0; i < N; i++ {
		if p[i] < box.Min[i] {
			x := p[i] - box.Min[i]
			dist += x * x
		}
		if p[i] > box.Max[i] {
			x := p[i] - box.Max[i]
			dist += x * x
		}
	}
	return math.Sqrt(dist)
}

// BoxNode is a node in a binary tree of boxes containing points.
type BoxNode struct {
	Box
	Mom  int
	Dau1 int
	Dau2 int
	Min  Point
	Max  Point
}

// KD-Tree
type KDTree struct {
	NumBoxes  int
	NumPoints int
	Points    []Point
	Boxes     []BoxNode
	PIdx      []int // index of points
	RIdx      []int // reverse index of points
}
