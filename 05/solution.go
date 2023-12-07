package puzzle05

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"advent2023/util/interval"
	"math"
	"slices"
	"strings"
)

type p struct {
	seeds    []int64
	mappings [7][]intervalMap
}

type intervalMap struct {
	interval.Interval
	Offset int64
}

func (p *p) Init(filename string) {
	blocks := util.GetBlocks(filename)

	p.seeds = util.ArrayStrToInt64(strings.Split(blocks[0][0], " ")[1:])

	for i, block := range blocks[1:] {
		var raw []intervalMap
		for _, line := range block[1:] {
			vals := util.ArrayStrToInt64(strings.SplitN(line, " ", 3))
			raw = append(raw, intervalMap{
				Interval: interval.Interval{Lower: vals[1], Upper: vals[1] + vals[2]},
				Offset:   vals[0] - vals[1],
			})
			slices.SortFunc(raw, func(l, r intervalMap) int { return l.Interval.Compare(r.Interval) })

			p.mappings[i] = make([]intervalMap, 0, len(raw))
			var start int64 = math.MinInt64
			for _, m := range raw {
				if start < m.Lower {
					p.mappings[i] = append(p.mappings[i], intervalMap{
						Interval: interval.Interval{Lower: start, Upper: m.Lower},
						Offset:   0,
					})
				}
				p.mappings[i] = append(p.mappings[i], m)
				start = m.Upper
			}
			if start < math.MaxInt64 {
				p.mappings[i] = append(p.mappings[i], intervalMap{
					Interval: interval.Interval{Lower: start, Upper: math.MaxInt64},
					Offset:   0,
				})
			}
		}
	}
}

func (p *p) doMap(x []interval.Interval, mapping []intervalMap) []interval.Interval {
	x = interval.Merge(x)
	mapped := make([]interval.Interval, 0)

	idxX, idxM := 0, 0
	for idxX < len(x) && idxM < len(mapping) {
		if x[idxX].Upper <= mapping[idxM].Upper {
			mapped = append(mapped, x[idxX].Shift(mapping[idxM].Offset))
			idxX++
			continue
		}

		if x[idxX].Lower >= mapping[idxM].Upper {
			idxM++
			continue
		}

		within := interval.Interval{Lower: x[idxX].Lower, Upper: mapping[idxM].Upper}
		mapped = append(mapped, within.Shift(mapping[idxM].Offset))

		x[idxX].Lower = mapping[idxM].Upper
		idxM++
	}

	return mapped
}

func (p *p) Solve1() any {
	seedIntervals := make([]interval.Interval, 0, len(p.seeds))
	for _, seed := range p.seeds {
		seedIntervals = append(seedIntervals, interval.Interval{Lower: seed, Upper: seed + 1})
	}

	for _, mapping := range p.mappings {
		seedIntervals = p.doMap(seedIntervals, mapping)
	}

	seedIntervals = interval.Merge(seedIntervals)
	return seedIntervals[0].Lower
}

func (p *p) Solve2() any {
	seedIntervals := make([]interval.Interval, 0, len(p.seeds)/2)
	for i, j := 0, 1; j < len(p.seeds); i, j = i+2, j+2 {
		seedIntervals = append(seedIntervals, interval.Interval{Lower: p.seeds[i], Upper: p.seeds[i] + p.seeds[j]})
	}
	for _, mapping := range p.mappings {
		seedIntervals = p.doMap(seedIntervals, mapping)
	}
	seedIntervals = interval.Merge(seedIntervals)
	return seedIntervals[0].Lower
}

func init() {
	puzzle.Register(5, &p{})
}
