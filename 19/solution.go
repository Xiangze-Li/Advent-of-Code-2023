package puzzle19

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"advent2023/util/interval"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type p struct {
	workflows map[string]workflow
	parts     [][4]int
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	reWorkflow := regexp.MustCompile(`^([[:alpha:]]+)\{(.+)\}$`)
	reRule := regexp.MustCompile(`^([xmas])([<>])(\d+):([[:alpha:]]+)$`)

	idx := 0
	p.workflows = make(map[string]workflow)
	for idx < len(lines) {
		if len(lines[idx]) == 0 {
			idx++
			break
		}

		m := reWorkflow.FindStringSubmatch(lines[idx])
		name := m[1]
		rules := strings.Split(m[2], ",")

		p.workflows[name] = workflow{
			name: name,
			rules: util.Reduce(rules[:len(rules)-1], func(now []rule, next string) []rule {
				mm := reRule.FindStringSubmatch(next)
				return append(now, rule{
					attr:   mm[1][0],
					op:     mm[2][0],
					val:    util.Must(strconv.Atoi(mm[3])),
					target: mm[4],
				})
			}, make([]rule, 0, len(rules)-1)),
			final: rules[len(rules)-1],
		}

		idx++
	}

	p.parts = make([][4]int, 0, len(lines)-idx)
	for idx < len(lines) {
		var x, m, a, s int
		util.Must(fmt.Sscanf(lines[idx], "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s))
		p.parts = append(p.parts, [4]int{x, m, a, s})
		idx++
	}
}

//nolint:gocognit // fvck off
func (p *p) Solve1() any {
	score := 0

	for _, part := range p.parts {
		cur := "in"
		for {
			wf := p.workflows[cur]
			ruleHit := false
			for _, rule := range wf.rules {
				attr := part[attrIdx[rule.attr]]
				if rule.op == '<' {
					if attr < rule.val {
						cur = rule.target
						ruleHit = true
					}
				} else {
					if attr > rule.val {
						cur = rule.target
						ruleHit = true
					}
				}
				if ruleHit {
					break
				}
			}
			if !ruleHit {
				cur = wf.final
			}
			if cur == "A" {
				score += part[0] + part[1] + part[2] + part[3]
				break
			}
			if cur == "R" {
				break
			}
		}
	}

	return score
}

func (p *p) calcCombination(cur string, accepted [4]interval.Interval) int64 {
	if cur == "R" {
		return 0
	}
	if cur == "A" {
		return util.Reduce(accepted[:], func(prod int64, next interval.Interval) int64 {
			return prod * (next.Upper - next.Lower)
		}, 1)
	}

	res := int64(0)
	wf := p.workflows[cur]
	for _, rule := range wf.rules {
		nextAcc := [4]interval.Interval{}
		copy(nextAcc[:], accepted[:])
		idx := attrIdx[rule.attr]

		if rule.op == '<' {
			nextAcc[idx].Upper = int64(rule.val)
			accepted[idx].Lower = int64(rule.val)
			res += p.calcCombination(rule.target, nextAcc)
		} else {
			nextAcc[idx].Lower = int64(rule.val + 1)
			accepted[idx].Upper = int64(rule.val + 1)
			res += p.calcCombination(rule.target, nextAcc)
		}
	}
	res += p.calcCombination(wf.final, accepted)
	return res
}

func (p *p) Solve2() any {
	const ub = 4001

	return p.calcCombination(
		"in",
		[4]interval.Interval{
			{Lower: 1, Upper: ub},
			{Lower: 1, Upper: ub},
			{Lower: 1, Upper: ub},
			{Lower: 1, Upper: ub},
		},
	)
}

func init() {
	puzzle.Register(19, &p{})
}
