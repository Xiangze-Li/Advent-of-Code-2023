package puzzle08

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"regexp"
)

type p struct {
	insts   []byte
	mapping map[string][2]string
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	re := regexp.MustCompile(`^(.{3}) = \((.{3}), (.{3})\)$`)

	p.insts = []byte(lines[0])
	p.mapping = make(map[string][2]string, len(lines)-2)

	for _, line := range lines[2:] {
		matches := re.FindStringSubmatch(line)
		p.mapping[matches[1]] = [2]string{matches[2], matches[3]}
	}
}

func (p *p) Solve1() any {
	const dst = "ZZZ"
	cur := "AAA"
	step := 0

	for cur != dst {
		inst := p.insts[step%len(p.insts)]
		if inst == 'L' {
			cur = p.mapping[cur][0]
		} else {
			cur = p.mapping[cur][1]
		}
		step++
	}

	return step
}

func isEnd(cur string) bool {
	return cur[2] == 'Z'
}

func isStart(cur string) bool {
	return cur[2] == 'A'
}

func (p *p) Solve2() any {
	start := []string{}

	for k := range p.mapping {
		if isStart(k) {
			start = append(start, k)
		}
	}

	steps := make([]int, len(start))
	for i := range steps {
		step := 0
		cur := start[i]
		for !isEnd(cur) {
			inst := p.insts[step%len(p.insts)]
			if inst == 'L' {
				cur = p.mapping[cur][0]
			} else {
				cur = p.mapping[cur][1]
			}
			step++
		}
		steps[i] = step
	}

	return util.LCM(steps[0], steps[1], steps[2:]...)
}

func init() {
	puzzle.Register(8, &p{})
}
