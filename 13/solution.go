//nolint:gocognit,cyclop,funlen // dont bother
package puzzle13

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"slices"
)

type p struct {
	grids [][][]byte
}

func (p *p) Init(filename string) {
	blocks := util.GetBlocks(filename)
	p.grids = make([][][]byte, 0, len(blocks))
	for _, block := range blocks {
		ongoing := make([][]byte, 0, len(block))
		for _, line := range block {
			ongoing = append(ongoing, []byte(line))
		}
		p.grids = append(p.grids, ongoing)
	}
}

func (p *p) Solve1() any {
	return util.Reduce(p.grids, func(acc int64, pattern [][]byte) int64 {
		lenI, lenJ := len(pattern), len(pattern[0])
		possI := 0
		possJs := make([]int, lenJ-1)
		for j := 1; j < lenJ; j++ {
			possJs[j-1] = j
		}

		for i := 0; i < lenI; i++ {
			if i > 0 { //nolint:nestif // dont bother
				if possI != 0 {
					mirroredI := possI*2 - i - 1
					if mirroredI >= 0 {
						if !slices.Equal(pattern[i], pattern[mirroredI]) {
							possI = 0
						}
					}
				}
				if possI == 0 {
					if slices.Equal(pattern[i], pattern[i-1]) {
						possI = i
					}
				}
			}

			stillPossJs := make([]int, 0, len(possJs))
			for _, possJ := range possJs {
				for delta := 0; delta < lenJ; delta++ {
					origJ, mirrJ := possJ+delta, possJ-1-delta
					if origJ >= lenJ || mirrJ < 0 {
						stillPossJs = append(stillPossJs, possJ)
						break
					}
					if pattern[i][origJ] != pattern[i][mirrJ] {
						break
					}
				}
			}
			possJs = stillPossJs
		}

		sum := 100 * int64(possI) //nolint:gomnd // dont bother
		if len(possJs) > 0 {
			sum += int64(possJs[0])
		}
		return acc + sum
	}, 0)
}

func (p *p) Solve2() any {
	var total int64

	for _, pattern := range p.grids {
		lenI, lenJ := len(pattern), len(pattern[0])
		possIs := make([]int, 0)
		diffIs := make([]int, 0)
		possJs := make([]int, lenJ-1)
		for j := 1; j < lenJ; j++ {
			possJs[j-1] = j
		}
		diffJs := make([]int, lenJ-1)

		for i := 0; i < lenI; i++ {
			if i > 0 { //nolint:nestif // dont bother
				validPossIs := make([]int, 0, len(possIs))
				validDiffIs := make([]int, 0, len(diffIs))
				if d := util.Diff(pattern[i], pattern[i-1]); d <= 1 {
					if !(d == 0 && i == lenI-1) {
						validPossIs = append(validPossIs, i)
						validDiffIs = append(validDiffIs, d)
					}
				}

				for ii := 0; ii < len(possIs); ii++ {
					possI, diffI := possIs[ii], diffIs[ii]
					mirroredI := possI*2 - i - 1
					if mirroredI < 0 {
						validPossIs = append(validPossIs, possI)
						validDiffIs = append(validDiffIs, diffI)
						continue
					}
					diffI += util.Diff(pattern[i], pattern[mirroredI])
					if (mirroredI == 0 && diffI == 1) ||
						(mirroredI > 0 && diffI <= 1) {
						validPossIs = append(validPossIs, possI)
						validDiffIs = append(validDiffIs, diffI)
					}
				}

				possIs, diffIs = validPossIs, validDiffIs
			}

			validPossJs := make([]int, 0, len(possJs))
			validDiffJs := make([]int, 0, len(diffJs))
			for jj := 0; jj < len(possJs); jj++ {
				possJ, diffJ := possJs[jj], diffJs[jj]
				for delta := 0; delta < lenJ; delta++ {
					origJ, mirrJ := possJ+delta, possJ-1-delta
					if origJ >= lenJ || mirrJ < 0 {
						validPossJs = append(validPossJs, possJ)
						validDiffJs = append(validDiffJs, diffJ)
						break
					}
					if pattern[i][origJ] != pattern[i][mirrJ] {
						diffJ++
						if diffJ > 1 {
							break
						}
					}
				}
			}
			possJs, diffJs = validPossJs, validDiffJs
		}

		var sum int
		for idx, possI := range possIs {
			if diffIs[idx] == 1 {
				sum += 100 * possI //nolint:gomnd // dont bother
			}
		}
		for idx, possJ := range possJs {
			if diffJs[idx] == 1 {
				sum += possJ
			}
		}
		total += int64(sum)
	}

	return total
}

func init() {
	puzzle.Register(13, &p{})
}
