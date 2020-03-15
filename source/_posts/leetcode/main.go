package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxPoints([][]int{
		{1, 1},
		{2, 1},
		{2, 2},
		{1, 4},
		{3, 3},
	}))

}

const (
	VerticalSlope   = "INF"
	HorizontalSlope = "NA"
)

type Point struct {
	X int
	Y int
	I int
}

func (p Point) ID() string {
	return fmt.Sprintf("%v:%v:%v", p.X, p.Y, p.I)
}

func (p Point) Vertical() Line {
	return Line{
		Points: map[string]Point{
			p.ID(): p,
		},
		Slope: VerticalSlope,
		C:     fmt.Sprint(p.X),
	}
}

func (p Point) Horizontal() Line {
	return Line{
		Points: map[string]Point{
			p.ID(): p,
		},
		Slope: HorizontalSlope,
		C:     fmt.Sprint(p.Y),
	}
}
func (p1 Point) Duplicate(p2 Point) bool {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return dx == 0 && dy == 0 && p1.I != p2.I
}

func (p1 Point) Oblique(p2 Point) Line {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	if dx == 0 {
		l := p1.Vertical()
		l.Add(p2)
		return l
	}

	if dy == 0 {
		l := p1.Horizontal()
		l.Add(p2)
		return l
	}

	d := gcd(dx, dy)
	if d > 1 || d < -1 {
		dy /= d
		dx /= d
	}
	slopeSymbol := dx/dy > 0
	slope := fmt.Sprintf("%v:%v:%v", slopeSymbol, abs(dx), abs(dy))

	c1 := p1.Y*dx - p1.X*dy
	c2 := dx
	cd := gcd(c1, c2)
	if cd > 1 || cd < -1 {
		c1 /= cd
		c2 /= cd
	}
	cSymbol := c1/c2 > 0

	c := fmt.Sprintf("%v:%v:%v", cSymbol, abs(c1), abs(c2))

	return Line{
		Points: map[string]Point{
			p1.ID(): p1,
			p2.ID(): p2,
		},
		Slope: slope,
		C:     c,
	}
}

type Line struct {
	Points map[string]Point
	Slope  string
	C      string
}

func (l Line) ID() string {
	return fmt.Sprintf("%s:%s", l.Slope, l.C)
}

func (l Line) Contain(p Point) bool {
	if _, ok := l.Points[p.ID()]; ok {
		return true
	} else {
		return false
	}
}

func (l Line) Add(p Point) {
	l.Points[p.ID()] = p
}

type Lines map[string]Line

func (ls Lines) Add(l Line) {
	if _, ok := ls[l.ID()]; ok {
		for _, p := range l.Points {
			ls[l.ID()].Add(p)
		}
	} else {
		ls[l.ID()] = l
	}
}

func maxPoints(points [][]int) int {
	ls := Lines{}
	var ps []Point
	for i, p := range points {
		ps = append(ps, Point{
			X: p[0],
			Y: p[1],
			I: i,
		})
	}

	for _, p1 := range ps {
		ls.Add(p1.Horizontal())
		ls.Add(p1.Vertical())

		duplicateSet := Line{
			Points: map[string]Point{},
		}
		tmpLines := Lines{}
		for _, p2 := range ps {
			if p1.Duplicate(p2) {
				duplicateSet.Add(p1)
				duplicateSet.Add(p2)
				continue
			}

			l := p1.Oblique(p2)
			tmpLines.Add(l)
		}

		for _, l := range tmpLines {
			ls.Add(l)
		}
	}

	var maxLine Line
	for _, l := range ls {
		if len(maxLine.Points) < len(l.Points) {
			maxLine = l
		}
	}
	return len(maxLine.Points)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
