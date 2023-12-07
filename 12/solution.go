package puzzle12

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"strings"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

//nolint:gocognit // this is algorithm.
func dp(row string, groups []int64) int64 {
	row += "."
	n, m := len(row), len(groups)
	mem := util.SliceND[int64](n+1, m+2, n+2).([][][]int64)
	mem[0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m+1; j++ {
			for k := 0; k < n+1; k++ {
				cur := mem[i][j][k]
				if cur == 0 {
					continue
				}
				if row[i] == '.' || row[i] == '?' {
					if k == 0 || k == int(groups[j-1]) {
						mem[i+1][j][0] += cur
					}
				}
				if row[i] == '#' || row[i] == '?' {
					if k == 0 {
						mem[i+1][j+1][k+1] += cur
					} else {
						mem[i+1][j][k+1] += cur
					}
				}
			}
		}
	}
	return mem[n][m][0]
}

func (p *p) Solve1() any {
	return util.Reduce(p.lines, func(acc int64, line string) int64 {
		sp := strings.SplitN(line, " ", 2)
		row := sp[0]
		groups := util.ArrayStrToInt64(strings.Split(sp[1], ","))
		return acc + dp(row, groups)
	}, 0)
}

func (p *p) Solve2() any {
	return util.Reduce(p.lines, func(acc int64, line string) int64 {
		sp := strings.SplitN(line, " ", 2)
		row := sp[0]
		row = strings.Repeat(row+"?", 4) + row
		groups := util.ArrayStrToInt64(strings.Split(sp[1], ","))
		groupsReal := make([]int64, len(groups)*5)
		for i := 0; i < 5; i++ {
			copy(groupsReal[i*len(groups):(i+1)*len(groups)], groups)
		}
		return acc + dp(row, groupsReal)
	}, 0)
}

func init() {
	puzzle.Register(12, &p{})
}
