package puzzle18

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
	"regexp"
	"strconv"
)

type p struct {
	dirs   []byte
	dists  []int64
	colors []string
}

func (p *p) Init(filename string) {
	re := regexp.MustCompile(`^([UDLR]) (\d+) \(#([0-9a-f]{6})\)$`)

	lines := util.GetLines(filename)
	p.dirs = make([]byte, 0, len(lines))
	p.dists = make([]int64, 0, len(lines))
	p.colors = make([]string, 0, len(lines))

	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		dir, dist, color := m[1][0], util.Must(strconv.ParseInt(m[2], 10, 64)), m[3]
		p.dirs = append(p.dirs, dir)
		p.dists = append(p.dists, dist)
		p.colors = append(p.colors, color)
	}
}

func (p *p) Solve1() any {
	var x, y int64
	var interior int64
	var perimeter int64

	for k := 0; k < len(p.dirs); k++ {
		dir, dist := p.dirs[k], p.dists[k]
		delta := c.Delta4[c.ConvertFromUDLR[dir]]
		xx, yy := x+int64(delta[0])*dist, y+int64(delta[1])*dist
		interior += x*yy - xx*y
		perimeter += dist
		x, y = xx, yy
	}

	if interior < 0 {
		interior = -interior
	}
	interior /= 2

	return interior + perimeter/2 + 1
}

func (p *p) Solve2() any {
	var directions = [...]byte{'R', 'D', 'L', 'U'}

	for i, color := range p.colors {
		b := []byte(color)
		p.dirs[i] = directions[b[5]-'0']
		p.dists[i] = util.Must(strconv.ParseInt(string(b[:5]), 16, 64))
	}
	return p.Solve1()
}

func init() {
	puzzle.Register(18, &p{})
}
