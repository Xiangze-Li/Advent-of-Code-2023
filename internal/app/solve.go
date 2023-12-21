package app

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func solve(day, part int, p puzzle.Puzzle, output io.Writer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(output, "error occured: %v\n", r)
		}
	}()
	fmt.Fprintf(output, "Solution for day %d part %d ", day, part)
	var solveFunc func() any
	if part == 1 {
		solveFunc = p.Solve1
	} else {
		solveFunc = p.Solve2
	}

	start := time.Now()
	res := solveFunc()
	used := time.Since(start)

	fmt.Fprintf(output, "in %s\n", used.String())
	if k := reflect.TypeOf(res).Kind(); k == reflect.Slice || k == reflect.Array {
		v := reflect.ValueOf(res)
		for i := 0; i < v.Len(); i++ {
			fmt.Fprintf(output, "\t%v\n", v.Index(i).Interface())
		}
	} else {
		fmt.Fprintf(output, "\t%v\n", res)
	}
}

func actionSolve(cCtx *cli.Context) error {
	args := cCtx.Args()
	if !args.Present() || args.Len() > 2 {
		cli.ShowSubcommandHelpAndExit(cCtx, 1)
	}
	var day, part int
	day = util.Must(strconv.Atoi(args.Get(0)))
	if args.Len() == 2 {
		part = util.Must(strconv.Atoi(args.Get(1)))
		if !(0 <= part && part <= 2) {
			return cli.Exit("Arugument `part` must be 0, 1 or 2. Got "+args.Get(1), 1)
		}
	}
	p, err := puzzle.Get(day)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	input := strings.ReplaceAll(cCtx.Path("input"), "{day}", fmt.Sprintf("%02d", day))

	st := time.Now()
	p.Init(input)
	used := time.Since(st)
	fmt.Fprintf(cCtx.App.Writer, "Solver for day %d initialized in %s\n", day, used.String())

	if part == 0 || part == 1 {
		solve(day, 1, p, cCtx.App.Writer)
	}
	if part == 0 || part == 2 {
		solve(day, 2, p, cCtx.App.Writer)
	}
	return nil
}

func cmdSolve() *cli.Command {
	return &cli.Command{
		Name:      "solve",
		Usage:     "Solve a puzzle",
		ArgsUsage: "day [part]",
		Action:    actionSolve,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "path to input file",
				Value:   "{day}/input.txt",
			},
		},
		HideHelp: true,
	}
}
