package puzzle23

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
	"maps"
)

type p struct {
	grid        [][]byte
	entry, exit [2]int
}

func (p *p) Init(filename string) {
	p.grid = util.GetGrid(filename)
	for j, c := range p.grid[0] {
		if c == '.' {
			p.entry = [2]int{0, j}
			break
		}
	}
	for j, c := range p.grid[len(p.grid)-1] {
		if c == '.' {
			p.exit = [2]int{len(p.grid) - 1, j}
			break
		}
	}
}

func (p *p) dfs1(cur [2]int, path map[[2]int]bool) int {
	if cur == p.exit {
		return 0
	}

	path[cur] = true
	maxium := len(path)
	for dir, delta := range c.Delta4 {
		next := [2]int{cur[0] + delta[0], cur[1] + delta[1]}
		if next[0] < 0 || next[0] >= len(p.grid) || next[1] < 0 || next[1] >= len(p.grid[0]) {
			continue
		}
		if p.grid[next[0]][next[1]] == '#' || path[next] {
			continue
		}
		switch p.grid[next[0]][next[1]] {
		case '.':
			maxium = max(maxium, p.dfs1(next, maps.Clone(path)))
		case 'v', '>', '^', '<':
			nextDir := mapArrow(p.grid[next[0]][next[1]])
			if nextDir == c.Opposite[dir] {
				break
			}
			pathC := maps.Clone(path)
			pathC[next] = true
			next = [2]int{next[0] + c.Delta4[nextDir][0], next[1] + c.Delta4[nextDir][1]}
			maxium = max(maxium, p.dfs1(next, pathC))
		}
	}
	return maxium
}

func mapArrow(ch byte) c.Direction {
	switch ch {
	case 'v':
		return c.S
	case '>':
		return c.E
	case '^':
		return c.N
	case '<':
		return c.W
	default:
		panic("invalid arrow")
	}
}

func (p *p) Solve1() any {
	return p.dfs1(p.entry, make(map[[2]int]bool))
}

func (p *p) buildAdj() [][]int {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	adj := util.SliceND[int](lenI*lenJ, lenI*lenJ).([][]int)
	flat := func(i, j int) int {
		return i*lenJ + j
	}

	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if p.grid[i][j] == '#' {
				continue
			}
			for _, delta := range c.Delta4 {
				ii, jj := i+delta[0], j+delta[1]
				if ii < 0 || ii >= lenI || jj < 0 || jj >= lenJ {
					continue
				}
				if p.grid[ii][jj] != '#' {
					adj[flat(i, j)][flat(ii, jj)] = -1
				}
			}
		}
	}

	return adj
}

func condense(adj [][]int) [][]int {
	l := len(adj)
	for i := 0; i < l; i++ {
		conn := []int{}
		for j := 0; j < l; j++ {
			if adj[i][j] != 0 {
				conn = append(conn, j)
			}
		}
		if len(conn) == 2 {
			a := adj[i][conn[0]]
			b := adj[i][conn[1]]
			adj[conn[0]][conn[1]] = a + b
			adj[conn[1]][conn[0]] = a + b
			adj[i][conn[0]] = 0
			adj[conn[0]][i] = 0
			adj[i][conn[1]] = 0
			adj[conn[1]][i] = 0
		}
	}
	return adj
}

func (p *p) dfs2(adj *[][]int, cur, exit int, path map[int]bool, length int, minium *int) {
	if cur == exit {
		*minium = min(*minium, length)
	}

	path[cur] = true
	for i, v := range (*adj)[cur] {
		if v == 0 || path[i] {
			continue
		}
		p.dfs2(adj, i, exit, maps.Clone(path), length+v, minium)
	}
}

func (p *p) Solve2() any {
	adj := condense(p.buildAdj())

	count := 0
	m := map[int]int{}
	for from, line := range adj {
		for _, v := range line {
			if v != 0 {
				m[from] = count
				count++
				break
			}
		}
	}

	shrinked := util.SliceND[int](count, count).([][]int)
	for from, newIdx := range m {
		for to, v := range adj[from] {
			if v != 0 {
				shrinked[newIdx][m[to]] = v
			}
		}
	}

	entry := m[p.entry[0]*len(p.grid[0])+p.entry[1]]
	exit := m[p.exit[0]*len(p.grid[0])+p.exit[1]]

	ans := 0
	p.dfs2(&shrinked, entry, exit, make(map[int]bool), 0, &ans)
	return -ans
}

func init() {
	puzzle.Register(23, &p{})
}
