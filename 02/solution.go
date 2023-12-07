package puzzle02

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type p struct {
	// lines []string
	games []game
}

type game struct {
	num   int64
	picks [][]colorPick
}

type colorPick struct {
	color string
	num   int64
}

func (p *p) Init(filename string) {
	reGame := regexp.MustCompile(`^Game (\d+): `)
	lines := util.GetLines(filename)
	p.games = make([]game, 0, len(lines))

	for _, line := range lines {
		g := game{}

		matchGame := reGame.FindStringSubmatchIndex(line)
		util.Assert(len(matchGame) == 4, "no game found")
		g.num = util.Must(strconv.ParseInt(line[matchGame[2]:matchGame[3]], 10, 64))

		for _, pick := range strings.Split(line[matchGame[1]:], ";") {
			colors := strings.Split(pick, ",")
			cp := make([]colorPick, 0, len(colors))
			for _, v := range colors {
				var num int64
				var color string
				fmt.Sscanf(strings.TrimSpace(v), "%d %s", &num, &color)
				cp = append(cp, colorPick{color, num})
			}
			g.picks = append(g.picks, cp)
		}

		p.games = append(p.games, g)
	}
}

func (p *p) Solve1() any {
	const r, g, b = 12, 13, 14

	return util.Reduce(p.games, func(acc int64, game game) int64 {
		for _, pick := range game.picks {
			for _, color := range pick {
				switch color.color {
				case "red":
					if color.num > r {
						return acc
					}
				case "green":
					if color.num > g {
						return acc
					}
				case "blue":
					if color.num > b {
						return acc
					}
				}
			}
		}
		return acc + game.num
	}, 0)
}

func (p *p) Solve2() any {
	return util.Reduce(p.games, func(acc int64, game game) int64 {
		var r, g, b int64

		for _, pick := range game.picks {
			for _, color := range pick {
				switch color.color {
				case "red":
					r = max(r, color.num)
				case "green":
					g = max(g, color.num)
				case "blue":
					b = max(b, color.num)
				}
			}
		}

		return acc + r*g*b
	}, 0)
}

func init() {
	puzzle.Register(2, &p{})
}
