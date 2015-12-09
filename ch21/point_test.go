package ch21_test

import (
	"math"
	"testing"

	"github.com/sbinet/go-numrecipes/ch21"
)

func TestPoint(t *testing.T) {
	for _, table := range []struct {
		p1   ch21.Point
		p2   ch21.Point
		dist float64
	}{
		{
			p1:   ch21.Point{},
			p2:   ch21.Point{},
			dist: 0.0,
		},
		{
			p1:   ch21.Point{1.0, 0.0},
			p2:   ch21.Point{0.0, 1.0},
			dist: math.Sqrt(2),
		},
	} {
		dist := ch21.Dist(table.p1, table.p2)
		if dist != table.dist {
			t.Fatalf(
				"error Point.Dist=%v want=%v. p1=%v, p2=%v\n",
				dist,
				table.dist,
				table.p1,
				table.p2,
			)
		}
	}
}

func TestBox(t *testing.T) {
	for _, table := range []struct {
		box  ch21.Box
		p    ch21.Point
		dist float64
	}{
		{
			box: ch21.Box{
				Min: ch21.Point{},
				Max: ch21.Point{},
			},
			p:    ch21.Point{},
			dist: 0,
		},
		{
			box: ch21.Box{
				Min: ch21.Point{},
				Max: ch21.Point{0, 1},
			},
			p:    ch21.Point{1, 1},
			dist: 1,
		},
		{
			box: ch21.Box{
				Min: ch21.Point{},
				Max: ch21.Point{0, 1},
			},
			p:    ch21.Point{1, 0.5},
			dist: 1,
		},
		{
			box: ch21.Box{
				Min: ch21.Point{0, 0},
				Max: ch21.Point{1, 1},
			},
			p:    ch21.Point{1, 1},
			dist: 0,
		},
		{
			box: ch21.Box{
				Min: ch21.Point{0, 0},
				Max: ch21.Point{1, 1},
			},
			p:    ch21.Point{-1, -1},
			dist: math.Sqrt(2),
		},
	} {
		dist := table.box.DistToPoint(table.p)
		if dist != table.dist {
			t.Fatalf(
				"error Box.Dist=%v want=%v. box=%v, point=%v\n",
				dist,
				table.dist,
				table.box,
				table.p,
			)
		}
	}
}
