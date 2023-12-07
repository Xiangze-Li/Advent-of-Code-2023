package puzzle09

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"strings"
)

type p struct {
	seqs [][]int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.seqs = make([][]int64, 0, len(lines))

	for _, line := range lines {
		p.seqs = append(p.seqs, util.ArrayStrToInt64(strings.Split(line, " ")))
	}
}

func calcDiffrencials(diffrencials *[][]int64, length, depth int) {
	diff := make([]int64, length-depth)
	allZero := true
	for i := 0; i < length-depth; i++ {
		diff[i] = (*diffrencials)[depth-1][i+1] - (*diffrencials)[depth-1][i]
		if diff[i] != 0 {
			allZero = false
		}
	}
	*diffrencials = append(*diffrencials, diff)
	if allZero {
		return
	}
	calcDiffrencials(diffrencials, length, depth+1)
}

func (p *p) Solve1() any {
	return util.Reduce(p.seqs, func(acc int64, seq []int64) int64 {
		length := len(seq)

		diffrencials := make([][]int64, 0, length)
		diffrencials = append(diffrencials, seq)
		calcDiffrencials(&diffrencials, length, 1)

		res := int64(0)
		for i := range diffrencials {
			res += diffrencials[i][len(diffrencials[i])-1]
		}
		return acc + res
	}, 0)
}

func (p *p) Solve2() any {
	return util.Reduce(p.seqs, func(acc int64, seq []int64) int64 {
		length := len(seq)

		diffrencials := make([][]int64, 0, length)
		diffrencials = append(diffrencials, seq)
		calcDiffrencials(&diffrencials, length, 1)

		res := int64(0)
		for i := len(diffrencials) - 1; i >= 0; i-- {
			res = diffrencials[i][0] - res
		}
		return acc + res
	}, 0)
}

func init() {
	puzzle.Register(9, &p{})
}
