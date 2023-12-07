package puzzle01

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
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() any {
	return util.Reduce(p.lines, func(acc int64, line string) int64 {
		first := strings.IndexAny(line, "0123456789")
		last := strings.LastIndexAny(line, "0123456789")
		util.Assert(first != -1 && last != -1, "no numbers found")
		return acc + util.Must(strconv.ParseInt(string([]byte{line[first], line[last]}), 10, 64))
	}, 0)
}

func (p *p) Solve2() any {
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[0-9])`)
	parse := func(s string) uint64 {
		switch s {
		case "one", "1":
			return 1
		case "two", "2":
			return 2
		case "three", "3":
			return 3
		case "four", "4":
			return 4
		case "five", "5":
			return 5
		case "six", "6":
			return 6
		case "seven", "7":
			return 7
		case "eight", "8":
			return 8
		case "nine", "9":
			return 9
		case "zero", "0":
			return 0
		default:
			panic("unreachable")
		}
	}

	return util.Reduce(p.lines, func(acc uint64, line string) uint64 {
		idx1 := re.FindStringIndex(line)
		util.Assert(len(idx1) == 2, "no numbers found")
		idx2 := []int{idx1[0], idx1[1]}
		for {
			temp := re.FindStringIndex(line[idx2[0]+1:])
			if len(temp) == 0 {
				break
			}
			idx2[0] += temp[0] + 1
			idx2[1] = idx2[0] + temp[1] - temp[0]
		}
		return acc + 10*parse(line[idx1[0]:idx1[1]]) + parse(line[idx2[0]:idx2[1]])
	}, 0)
}

func init() {
	puzzle.Register(1, &p{})
}
