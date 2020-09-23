package main

import (
	"fmt"
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
	case "go":
		fmt.Println("Generate file go")
		pg := generator.NewGo(arg)
		pg.Generate()
		break
	default:
		df.Default()
	}
}
