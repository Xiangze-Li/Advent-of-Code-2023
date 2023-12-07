package puzzle04

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"cmp"
	"slices"
	"strings"
)

type p struct {
	cards []card
}

type card struct {
	winning []int64
	having  []int64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.cards = make([]card, 0, len(lines))

	for _, line := range lines {
		splitCard := strings.SplitN(line, ":", 2)
		util.Assert(len(splitCard) == 2, "invalid input")
		split := strings.SplitN(splitCard[1], "|", 2)
		util.Assert(len(split) == 2, "invalid input")

		winnings := util.ArrayStrToInt64(strings.Fields(split[0]))
		havings := util.ArrayStrToInt64(strings.Fields(split[1]))
		slices.Sort(winnings)
		slices.Sort(havings)

		p.cards = append(p.cards, card{winnings, havings})
	}
}

func (p *p) handleCard(card card) int {
	var w, h int
	var count int
	for w < len(card.winning) && h < len(card.having) {
		switch cmp.Compare(card.winning[w], card.having[h]) {
		case -1:
			w++
		case 0:
			count++
			h++
		case 1:
			h++
		}
	}
	return count
}

func (p *p) Solve1() any {
	return util.Reduce(p.cards, func(acc int64, card card) int64 {
		count := p.handleCard(card)
		if count == 0 {
			return acc
		}
		return acc + int64(1)<<count
	}, 0)
}

func (p *p) Solve2() any {
	var counts = make([]int64, len(p.cards))
	for i, card := range p.cards {
		counts[i]++
		count := p.handleCard(card)
		for j := i + 1; j < len(p.cards) && j <= i+count; j++ {
			counts[j] += counts[i]
		}
	}

	return util.Reduce(counts, func(acc, next int64) int64 { return acc + next }, 0)
}

func init() {
	puzzle.Register(4, &p{})
}
