package app

import (
	"github.com/urfave/cli/v2"
)

var App *cli.App //nolint:gochecknoglobals // global registry

func init() {
	App = cli.NewApp()
	App.Usage = "Solver for Advent of Code 2023"
	App.Authors = []*cli.Author{{Name: "Xiangze Li", Email: "lee_johnson@qq.com"}}
	App.Version = "0.1.0"
	App.HideHelp = true
	App.HideVersion = true
	App.Commands = []*cli.Command{
		cmdPrepare(),
		cmdSolve(),
	}
}
