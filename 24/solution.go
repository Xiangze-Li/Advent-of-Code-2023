package puzzle24

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
)

const (
	lb = 200_000_000_000_000
	ub = 400_000_000_000_000
)

type p struct {
	points     [][3]float64
	velocities [][3]float64

	isExample bool
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.points = make([][3]float64, 0, len(lines))
	p.velocities = make([][3]float64, 0, len(lines))

	for _, line := range lines {
		sp := strings.SplitN(line, "@", 2)

		point := util.ArrayStrToFloat64(strings.SplitN(sp[0], ",", 3))
		velocity := util.ArrayStrToFloat64(strings.SplitN(sp[1], ",", 3))

		p.points = append(p.points, [3]float64{point[0], point[1], point[2]})
		p.velocities = append(p.velocities, [3]float64{velocity[0], velocity[1], velocity[2]})
	}

	if filename == "24/example.txt" {
		p.isExample = true
	}
}

func intersect2D(p1, v1, p2, v2 [3]float64) (bool, [2]float64) {
	vCross := v1[0]*v2[1] - v1[1]*v2[0]
	if vCross == 0 {
		return false, [2]float64{}
	}

	t2 := (v1[1]*(p1[0]-p2[0]) - v1[0]*(p1[1]-p2[1])) / (-vCross)
	if t2 < 0 {
		return false, [2]float64{}
	}
	t1 := (v2[0]*(p1[1]-p2[1]) - v2[1]*(p1[0]-p2[0])) / (vCross)
	if t1 < 0 {
		return false, [2]float64{}
	}

	pIn := [2]float64{}
	pIn[0] = p2[0] + v2[0]*t2
	pIn[1] = p2[1] + v2[1]*t2

	return true, pIn
}

func (p *p) Solve1() any {
	var count int
	var lb, ub float64 = lb, ub
	if p.isExample {
		lb, ub = 7, 27
	}

	for i := 0; i < len(p.points); i++ {
		for j := i + 1; j < len(p.points); j++ {
			can, pIn := intersect2D(p.points[i], p.velocities[i], p.points[j], p.velocities[j])
			if can && pIn[0] >= lb && pIn[0] <= ub && pIn[1] >= lb && pIn[1] <= ub {
				count++
			}
		}
	}

	return count
}

func (p *p) Solve2() any {
	// For the math, see 24/note.md

	a := mat.NewDense(6, 6, nil)
	b := mat.NewDense(6, 1, nil)
	x := mat.NewDense(6, 1, nil)

	for i, j := 0, 1; j <= 2; j++ {
		vDiff := sub(p.velocities[i], p.velocities[j])
		pDiff := sub(p.points[i], p.points[j])
		crossI, crossJ := cross(p.points[i], p.velocities[i]), cross(p.points[j], p.velocities[j])
		bb := sub(crossJ, crossI)

		a.SetRow(3*(j-1)+0, []float64{0, -vDiff[2], vDiff[1], 0, pDiff[2], -pDiff[1]})
		a.SetRow(3*(j-1)+1, []float64{vDiff[2], 0, -vDiff[0], -pDiff[2], 0, pDiff[0]})
		a.SetRow(3*(j-1)+2, []float64{-vDiff[1], vDiff[0], 0, pDiff[1], -pDiff[0], 0})
		b.Set(3*(j-1)+0, 0, bb[0])
		b.Set(3*(j-1)+1, 0, bb[1])
		b.Set(3*(j-1)+2, 0, bb[2])
	}

	util.Must(0, x.Solve(a, b))

	var sum float64
	for i := 0; i < 3; i++ {
		sum += math.Round(x.At(i, 0))
	}
	return int(sum)
}

func init() {
	puzzle.Register(24, &p{})
}

func sub(a, b [3]float64) [3]float64 {
	return [3]float64{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func cross(a, b [3]float64) [3]float64 {
	return [3]float64{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}
