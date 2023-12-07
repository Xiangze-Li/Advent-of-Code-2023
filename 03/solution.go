package puzzle03

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
	"bytes"
	"strconv"
)

type p struct {
	grids [][]byte
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	lenI, lenJ := len(lines), len(lines[0])
	p.grids = make([][]byte, lenI+2)

	p.grids[0] = bytes.Repeat([]byte{'.'}, lenJ+2)
	for i, line := range lines {
		p.grids[i+1] = make([]byte, lenJ+2)
		p.grids[i+1][0] = '.'
		copy(p.grids[i+1][1:lenJ+1], []byte(line))
		p.grids[i+1][lenJ+1] = '.'
	}
	p.grids[lenI+1] = bytes.Repeat([]byte{'.'}, lenJ+2)
}

func (p *p) isDigit(i, j int) bool {
	return p.grids[i][j] >= '0' && p.grids[i][j] <= '9'
}

func (p *p) isPart(i, j int) bool {
	return p.grids[i][j] != '.' && !p.isDigit(i, j)
}

func (p *p) expandNumber(i, j int, vis *[][]bool) int64 {
	ongoing := []byte{}

	if (*vis)[i][j] || !p.isDigit(i, j) {
		return 0
	}

	ongoing = append(ongoing, p.grids[i][j])
	(*vis)[i][j] = true

	for k, delta := range [2][2]int{{0, -1}, {0, 1}} {
		ii, jj := i+delta[0], j+delta[1]
		for !(*vis)[ii][jj] && p.isDigit(ii, jj) {
			if k == 0 {
				ongoing = append([]byte{p.grids[ii][jj]}, ongoing...)
			} else {
				ongoing = append(ongoing, p.grids[ii][jj])
			}
			(*vis)[ii][jj] = true
			ii, jj = ii+delta[0], jj+delta[1]
		}
	}

	return util.Must(strconv.ParseInt(string(ongoing), 10, 64))
}

func (p *p) Solve1() any {
	var sum int64
	var lenI, lenJ = len(p.grids), len(p.grids[0])
	var vis = util.SliceND[bool](lenI, lenJ).([][]bool)

	for i := 1; i < lenI-1; i++ {
		for j := 1; j < lenJ-1; j++ {
			if p.isPart(i, j) {
				for _, delta := range c.Delta8 {
					ii, jj := i+delta[0], j+delta[1]
					if !vis[ii][jj] && p.isDigit(ii, jj) {
						sum += p.expandNumber(ii, jj, &vis)
					}
				}
			}
		}
	}

	return sum
}

func (p *p) Solve2() any {
	var sum int64
	var lenI, lenJ = len(p.grids), len(p.grids[0])
	var vis = util.SliceND[bool](lenI, lenJ).([][]bool)

	for i := 1; i < lenI-1; i++ {
		for j := 1; j < lenJ-1; j++ {
			if p.grids[i][j] == '*' {
				var count, prod int64 = 0, 1
				for _, delta := range c.Delta8 {
					ii, jj := i+delta[0], j+delta[1]
					if !vis[ii][jj] && p.isDigit(ii, jj) {
						prod *= p.expandNumber(ii, jj, &vis)
						count++
					}
				}
				if count == 2 {
					sum += prod
				}
			}
		}
	}

	return sum
}

func init() {
	puzzle.Register(3, &p{})
}
