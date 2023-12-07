package puzzle07

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"fmt"
	"sort"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() any {
	hands := make([]hand, 0, len(p.lines))
	for _, line := range p.lines {
		var c string
		var bid uint64
		util.Must(fmt.Sscan(line, &c, &bid))

		h := hand{
			Cards: convertCardStrength1(c),
			Bid:   bid,
		}
		h.parseType1()
		hands = append(hands, h)
	}

	sort.Sort(handSlice(hands))
	return util.ReduceIndex(hands, func(acc uint64, idx int, hand hand) uint64 {
		return acc + hand.Bid*uint64(idx+1)
	}, 0)
}

func (p *p) Solve2() any {
	hands := make([]hand, 0, len(p.lines))
	for _, line := range p.lines {
		var c string
		var bid uint64
		util.Must(fmt.Sscan(line, &c, &bid))

		h := hand{
			Cards: convertCardStrength2(c),
			Bid:   bid,
		}
		h.parseType2()
		hands = append(hands, h)
	}

	sort.Sort(handSlice(hands))
	return util.ReduceIndex(hands, func(acc uint64, idx int, hand hand) uint64 {
		return acc + hand.Bid*uint64(idx+1)
	}, 0)
}

func init() {
	puzzle.Register(7, &p{})
}
