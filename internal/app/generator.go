package app

import (
	"advent2023/util"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/urfave/cli/v2"
)

func actionPrepare(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}
	day, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return fmt.Errorf("invalid day: %w", err)
	}
	if day < 1 || day > 25 {
		return fmt.Errorf("invalid day: %d, should between 1 and 25", day)
	}

	dayName := fmt.Sprintf("%02d", day)
	if err = os.Mkdir(dayName, 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(c.App.ErrWriter, "error occured when downloading input: %v\n", err)
			}
		}()
		//nolint:gosec // checked input
		cmd := exec.Command("aoc",
			"download", "-I", "-d", strconv.Itoa(day), "-y2023", "-i", dayName+"/input.txt")
		util.Must(0, cmd.Run())
		util.Must(os.OpenFile(dayName+"/example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)).Close()
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(c.App.ErrWriter, "error occured when generating solver: %v\n", err)
			}
		}()
		file := util.Must(os.OpenFile(dayName+"/solution.go", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644))
		defer file.Close()
		data := tmplSolver
		data = strings.ReplaceAll(data, "{{.DayString}}", dayName)
		data = strings.ReplaceAll(data, "{{.Day}}", strconv.Itoa(day))
		util.Must(file.WriteString(data))
	}()

	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(c.App.ErrWriter, "error occured when generating registery: %v\n", err)
			}
		}()

		data := tmplRegistry
		data = strings.ReplaceAll(data, "{{.DayString}}", dayName)
		data = strings.ReplaceAll(data, "{{.Day}}", strconv.Itoa(day))
		_ = os.MkdirAll("internal/registry", 0755)
		file := util.Must(os.OpenFile("internal/registry/"+dayName+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644))
		defer file.Close()
		util.Must(file.WriteString(data))
	}()

	wg.Wait()
	return nil
}

func cmdPrepare() *cli.Command {
	return &cli.Command{
		Name:      "prepare",
		Usage:     "Prepare for a new day",
		ArgsUsage: "day",
		Action:    actionPrepare,
		HideHelp:  true,
	}
}

const tmplSolver = `package puzzle{{.DayString}}

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
)

type p struct {
	lines []string
}

func (p *p) Init(filename string) {
	p.lines = util.GetLines(filename)
}

func (p *p) Solve1() any {
	return 0
}

func (p *p) Solve2() any {
	return 0
}

func init() {
	puzzle.Register({{.Day}}, &p{})
}
`

const tmplRegistry = `package reg

import _ "advent2023/{{.DayString}}" // register puzzle solver
`
