package main

import (
	"os"

	"github.com/mmuflih/file-generator/generator"
)

func main() {
	df := generator.NewDefault()
	arg := os.Args[1:]
	if len(arg) < 1 {
		df.Default()
		return
	}
	switch arg[0] {
	case "php":
		pg := generator.NewPhpGo(arg)
		pg.Generate()
		break
	default:
		df.Default()
	}
}
