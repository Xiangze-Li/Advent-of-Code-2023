package puzzle20

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type p struct {
	moduleType map[string]modType
	to         map[string][]string
	from       map[string][]string
	moduleIdx  map[string]int
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.moduleType = make(map[string]modType, len(lines))
	p.moduleIdx = make(map[string]int, len(lines))
	p.to = make(map[string][]string, len(lines))
	p.from = make(map[string][]string, len(lines))

	for _, line := range lines {
		sp := strings.SplitN(line, " -> ", 2)
		name := sp[0]
		dests := strings.Split(sp[1], ", ")

		switch {
		case name == "broadcaster":
			p.moduleType[name] = broadcaster
		case name[0] == '%':
			name = name[1:]
			p.moduleType[name] = flipflop
		case name[0] == '&':
			name = name[1:]
			p.moduleType[name] = conjunction
		default:
			panic("unknown module type")
		}

		p.to[name] = dests
		for _, dest := range dests {
			p.from[dest] = append(p.from[dest], name)
		}
	}

	modCount := 0
	for name := range p.moduleType {
		p.moduleIdx[name] = modCount
		modCount++
	}
}

//nolint:gocognit
func (p *p) Solve1() any {
	const target = 1000
	modCount := len(p.moduleIdx)
	mem := util.SliceND[byte](modCount, modCount).([][]byte)
	vis := map[string][3]int{hash(mem): {0, 0, 0}}

	cycleFound := false
	lowCount, highCount := 0, 0

	for round := 1; round <= target; round++ {
		lowCount++ // button -> broadcaster
		queue := []pulse{}
		for _, to := range p.to["broadcaster"] {
			queue = append(queue, pulse{"broadcaster", to, 0})
		}
		for len(queue) > 0 {
			pul := queue[0]
			queue = queue[1:]

			if pul.level != 0 {
				highCount++
			} else {
				lowCount++
			}

			switch p.moduleType[pul.to] {
			case dumb, broadcaster: // noop
			case flipflop:
				if pul.level != 0 {
					continue
				}
				id := p.moduleIdx[pul.to]
				if mem[id][0] == 0 {
					mem[id][0] = 1
					for _, to := range p.to[pul.to] {
						queue = append(queue, pulse{pul.to, to, 1})
					}
				} else {
					mem[id][0] = 0
					for _, to := range p.to[pul.to] {
						queue = append(queue, pulse{pul.to, to, 0})
					}
				}
			case conjunction:
				idFrom, idTo := p.moduleIdx[pul.from], p.moduleIdx[pul.to]
				mem[idTo][idFrom] = pul.level
				outLevel := byte(0)
				for _, v := range p.from[pul.to] {
					if mem[idTo][p.moduleIdx[v]] == 0 {
						outLevel = 1
						break
					}
				}
				for _, to := range p.to[pul.to] {
					queue = append(queue, pulse{pul.to, to, outLevel})
				}
			}
		}
		if !cycleFound {
			h := hash(mem)
			if seen, ok := vis[h]; ok {
				cycleLength := round - seen[0]
				lowPerCycle := lowCount - seen[1]
				highPerCycle := highCount - seen[2]

				fullCycle := (target - round) / cycleLength
				lowCount += lowPerCycle * fullCycle
				highCount += highPerCycle * fullCycle
				round += cycleLength*fullCycle + 1

				cycleFound = true
			}
			vis[h] = [3]int{round, lowCount, highCount}
		}
	}

	return lowCount * highCount
}

func (p *p) Solve2() any {
	return "brute-force will not work"
}

func init() {
	puzzle.Register(20, &p{})
}

func hash(mem [][]byte) string {
	h := md5.Sum(bytes.Join(mem, nil))
	return hex.EncodeToString(h[:])
}
