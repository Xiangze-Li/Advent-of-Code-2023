package puzzle11

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
)

type p struct {
	countI, countJ []int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)

	p.countI = make([]int64, len(lines))
	p.countJ = make([]int64, len(lines[0]))

	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				p.countI[i]++
				p.countJ[j]++
			}
		}
	}
}

func (p *p) calculate(expansion int64) int64 {
	var sum, ongoing, counting int64
	for _, count := range p.countI {
		if count == 0 {
			ongoing += counting * expansion
		} else {
			ongoing += counting
			sum += ongoing * count
			counting += count
		}
	}
	ongoing, counting = 0, 0
	for _, count := range p.countJ {
		if count == 0 {
			ongoing += counting * expansion
		} else {
			ongoing += counting
			sum += ongoing * count
			counting += count
		}
	}
	return sum
}

func (p *p) Solve1() any {
	return p.calculate(2)
}

func (p *p) Solve2() any {
	const expansion = 1_000_000
	return p.calculate(expansion)
}

func init() {
	puzzle.Register(11, &p{})
}
