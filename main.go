package main

import (
	"advent2023/internal/app"
	"fmt"
	"os"
)

func main() {
	if err := app.App.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
