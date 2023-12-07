package puzzle25

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"strings"

	"github.com/dominikbraun/graph"
)

type p struct {
	// nodes []string
	g graph.Graph[string, string]
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	nodes := map[string]bool{}
	p.g = graph.New(graph.StringHash)
	// p.nodeIdx = make(map[string]int)
	// p.adj = make(map[int]map[int]bool)

	for _, line := range lines {
		sp := strings.Split(line, ":")
		from := strings.TrimSpace(sp[0])
		to := strings.Fields(sp[1])
		if !nodes[from] {
			nodes[from] = true
			p.g.AddVertex(from)
		}
		for _, t := range to {
			if !nodes[t] {
				nodes[t] = true
				p.g.AddVertex(t)
			}
			p.g.AddEdge(from, t)
		}
	}
}

func (p *p) Solve1() any {
	return "DONT WORK YET. LIFE IS SHORT, JUST USE PYTHON."
}

func (p *p) Solve2() any {
	return "THANK YOU MR.WASTL!"
}

func init() {
	puzzle.Register(25, &p{})
}
