package puzzle06

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"math"
	"strconv"
	"strings"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) solveGame(t, d int64) int64 {
	util.Assert(t*t >= 4*d, "no solution")

	sqrtD := math.Sqrt(float64(t*t - 4*d))
	x1 := (float64(t) - sqrtD) / 2
	x2 := (float64(t) + sqrtD) / 2

	mini := math.Ceil(x1)
	if mini == x1 {
		mini++
	}
	maxi := math.Floor(x2)
	if maxi == x2 {
		maxi--
	}

	return int64(maxi) - int64(mini) + 1
}

func (p *p) Solve1() any {
	var prod int64 = 1

	times := util.ArrayStrToInt64(strings.Fields(p.lines[0])[1:])
	dists := util.ArrayStrToInt64(strings.Fields(p.lines[1])[1:])

	for i := 0; i < len(times); i++ {
		t, d := times[i], dists[i]
		prod *= p.solveGame(t, d)
	}

	return prod
}

func (p *p) Solve2() any {
	t := util.Must(strconv.ParseInt((strings.Join(strings.Fields(p.lines[0])[1:], "")), 10, 64))
	d := util.Must(strconv.ParseInt((strings.Join(strings.Fields(p.lines[1])[1:], "")), 10, 64))

	return p.solveGame(t, d)
}

func init() {
	puzzle.Register(6, &p{})
}
