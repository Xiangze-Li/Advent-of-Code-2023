package puzzle22

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"maps"
	"strings"
)

type p struct {
	cubes []cube
	exist map[[3]int]int
}

type cube struct {
	lower, upper [3]int
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)

	p.cubes = make([]cube, 0, len(lines))
	p.exist = make(map[[3]int]int)

	for idx, line := range lines {
		sp := strings.Split(line, "~")

		lo := util.ArrayStrToInt64(strings.Split(sp[0], ","))
		hi := util.ArrayStrToInt64(strings.Split(sp[1], ","))
		p.cubes = append(p.cubes, cube{
			lower: [3]int{int(lo[0]), int(lo[1]), int(lo[2])},
			upper: [3]int{int(hi[0]), int(hi[1]), int(hi[2])},
		})

		for x := lo[0]; x <= hi[0]; x++ {
			for y := lo[1]; y <= hi[1]; y++ {
				for z := lo[2]; z <= hi[2]; z++ {
					p.exist[[3]int{int(x), int(y), int(z)}] = idx + 1
				}
			}
		}
	}

	p.fall()
}

func (p *p) fall() {
	for {
		moved := false

		for idx, c := range p.cubes {
			canMove := 0
			for x := c.lower[0]; x <= c.upper[0]; x++ {
				for y := c.lower[1]; y <= c.upper[1]; y++ {
					var z = c.lower[2] - 1
					for ; z > canMove; z-- {
						if p.exist[[3]int{x, y, z}] != 0 {
							break
						}
					}
					canMove = max(canMove, z)
				}
			}
			canMove++
			if canMove < c.lower[2] {
				moved = true
				p.moveCube(idx, canMove)
			}
		}

		if !moved {
			break
		}
	}

	maps.DeleteFunc(p.exist, func(_ [3]int, v int) bool { return v == 0 })
}

func (p *p) moveCube(idx int, canMove int) {
	c := p.cubes[idx]
	delta := c.lower[2] - canMove
	height := c.upper[2] - c.lower[2] + 1

	for x := c.lower[0]; x <= c.upper[0]; x++ {
		for y := c.lower[1]; y <= c.upper[1]; y++ {
			for z := canMove; z <= canMove+height-1; z++ {
				p.exist[[3]int{x, y, z}] = idx + 1
				p.exist[[3]int{x, y, z + delta}] = 0
			}
		}
	}

	c.lower[2] = canMove
	c.upper[2] = canMove + height - 1
	p.cubes[idx] = c
}

func (p *p) Solve1() any {
	supporting := map[int]map[int]bool{}
	supported := map[int]int{}

	for idx, c := range p.cubes {
		ing := map[int]bool{}

		for x := c.lower[0]; x <= c.upper[0]; x++ {
			for y := c.lower[1]; y <= c.upper[1]; y++ {
				zHigh := c.upper[2] + 1
				if cc := p.exist[[3]int{x, y, zHigh}]; cc != 0 {
					ing[cc] = true
				}
			}
		}

		supporting[idx+1] = ing
		for k := range ing {
			supported[k]++
		}
	}

	res := 0
	for idx := range p.cubes {
		canBreak := true
		for ed := range supporting[idx+1] {
			if supported[ed] == 1 {
				canBreak = false
				break
			}
		}
		if canBreak {
			res++
		}
	}
	return res
}

func (p *p) buildAdj() [][]int {
	nCube := len(p.cubes)
	adj := util.SliceND[int](nCube+1, nCube+1).([][]int)

	for idx, c := range p.cubes {
		for x := c.lower[0]; x <= c.upper[0]; x++ {
			for y := c.lower[1]; y <= c.upper[1]; y++ {
				zHigh := c.upper[2] + 1
				if cc := p.exist[[3]int{x, y, zHigh}]; cc != 0 {
					adj[idx+1][cc] = 1
					adj[cc][idx+1] = -1
				}
			}
		}
	}
	return adj
}

func (p *p) Solve2() any {
	adj := p.buildAdj()
	res := 0

	for i := range p.cubes {
		willFall := map[int]bool{i + 1: true}
		q := []int{i + 1}

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			for j, s := range adj[cur] {
				if s == 1 {
					thisWillFall := true
					for k, ss := range adj[j] {
						if ss == -1 && !willFall[k] {
							thisWillFall = false
							break
						}
					}
					if thisWillFall {
						willFall[j] = true
						q = append(q, j)
					}
				}
			}
		}

		res += len(willFall) - 1
	}

	return res
}

func init() {
	puzzle.Register(22, &p{})
}
