package puzzle14

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"bytes"
)

type p struct {
	grid [][]byte
}

func (p *p) Init(filename string) {
	p.grid = util.GetGrid(filename)
}

func (p *p) Solve1() any {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	var sum uint64

	for j := 0; j < lenJ; j++ {
		blockI := -1
		ongoing := 0
		for i := 0; i < lenI; i++ {
			switch p.grid[i][j] {
			case '.':
				continue
			case '#':
				sum += uint64(ongoing*(lenI-blockI-1) - ongoing*(ongoing-1)/2)
				ongoing = 0
				blockI = i
			case 'O':
				ongoing++
			}
		}
		if ongoing > 0 {
			sum += uint64(ongoing*(lenI-blockI-1) - ongoing*(ongoing-1)/2)
		}
	}

	return sum
}

func (p *p) Solve2() any {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	const billion = 1_000_000_000
	breakpoint := billion
	seen := map[string]int{}

	for k := 0; k < billion; k++ {
		p.cycle()

		if breakpoint == billion {
			h := string(bytes.Join(p.grid, nil))
			if i, ok := seen[h]; ok {
				breakpoint = k + (billion-i)%(k-i) - 1
			} else {
				seen[h] = k
			}
		}

		if k == breakpoint {
			break
		}
	}

	var sum uint64
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if p.grid[i][j] == 'O' {
				sum += uint64(lenI - i)
			}
		}
	}

	return sum
}

func init() {
	puzzle.Register(14, &p{})
}

//nolint:funlen,gocognit // dont bother
func (p *p) cycle() {
	lenI, lenJ := len(p.grid), len(p.grid[0])

	for j := 0; j < lenJ; j++ {
		blockI := -1
		ongoing := 0
		for i := 0; i < lenI; i++ {
			switch p.grid[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < ongoing; n++ {
					p.grid[blockI+n+1][j] = 'O'
				}
				ongoing = 0
				blockI = i
			case 'O':
				p.grid[i][j] = '.'
				ongoing++
			}
		}
		for n := 0; n < ongoing; n++ {
			p.grid[blockI+n+1][j] = 'O'
		}
	}

	for i := 0; i < lenI; i++ {
		blockJ := -1
		ongoing := 0
		for j := 0; j < lenJ; j++ {
			switch p.grid[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < ongoing; n++ {
					p.grid[i][blockJ+n+1] = 'O'
				}
				ongoing = 0
				blockJ = j
			case 'O':
				p.grid[i][j] = '.'
				ongoing++
			}
		}
		for n := 0; n < ongoing; n++ {
			p.grid[i][blockJ+n+1] = 'O'
		}
	}

	for j := 0; j < lenJ; j++ {
		blockI := lenI
		ongoing := 0
		for i := lenI - 1; i >= 0; i-- {
			switch p.grid[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < ongoing; n++ {
					p.grid[blockI-n-1][j] = 'O'
				}
				ongoing = 0
				blockI = i
			case 'O':
				p.grid[i][j] = '.'
				ongoing++
			}
		}
		for n := 0; n < ongoing; n++ {
			p.grid[blockI-n-1][j] = 'O'
		}
	}

	for i := 0; i < lenI; i++ {
		blockJ := lenJ
		ongoing := 0
		for j := lenJ - 1; j >= 0; j-- {
			switch p.grid[i][j] {
			case '.':
				continue
			case '#':
				for n := 0; n < ongoing; n++ {
					p.grid[i][blockJ-n-1] = 'O'
				}
				ongoing = 0
				blockJ = j
			case 'O':
				p.grid[i][j] = '.'
				ongoing++
			}
		}
		for n := 0; n < ongoing; n++ {
			p.grid[i][blockJ-n-1] = 'O'
		}
	}
}
