package puzzle15

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"regexp"
	"strconv"
	"strings"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = strings.Split(util.GetLines(filename)[0], ",")
}

func hash(line string) int64 {
	cur := int64(0)
	for _, b := range []byte(line) {
		cur += int64(b)
		cur *= 17
		cur %= 256
	}
	return cur
}

func (p *p) Solve1() any {
	return util.Reduce(p.lines, func(acc int64, line string) int64 {
		return acc + hash(line)
	}, 0)
}

func (p *p) Solve2() any {
	type lens struct {
		Label string
		Focal int
	}
	boxes := [256][]lens{}

	re := regexp.MustCompile(`^([a-z]+)([=\-])([1-9]?)$`)

	for _, line := range p.lines {
		m := re.FindStringSubmatch(line)
		label, op, f := m[1], m[2], m[3]
		h := int(hash(label))
		if op == "=" {
			focal := util.Must(strconv.Atoi(f))
			i := 0
			for i < len(boxes[h]) {
				if boxes[h][i].Label == label {
					boxes[h][i].Focal = focal
					break
				}
				i++
			}
			if i == len(boxes[h]) {
				boxes[h] = append(boxes[h], lens{label, focal})
			}
		} else {
			rem := []lens{}
			for _, l := range boxes[h] {
				if l.Label != label {
					rem = append(rem, l)
				}
			}
			boxes[h] = rem
		}
	}

	var sum uint64
	for iBox, box := range boxes {
		for iSlot, lens := range box {
			sum += uint64((iBox + 1) * lens.Focal * (iSlot + 1))
		}
	}
	return sum
}

func init() {
	puzzle.Register(15, &p{})
}
