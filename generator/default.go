package generator

import "fmt"

type default0 struct {
}

type Default0 interface {
	Default()
}

func NewDefault() Default0 {
	return &default0{}
}

func (df default0) Default() {
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("    fg command type [arguments]")
	fmt.Println()
	fmt.Println("AVAILABLE COMMANDS")
	fmt.Println("\n" +
		"    php      To Generate PHP File" +
		"\n")
	fmt.Println("AVAILABLE TYPES")
	fmt.Println("\n" +
		"    class    To Generate class" +
		"\n")
	fmt.Println("EXAMPLE")
	fmt.Println("\n" +
		"    fg php class App/Core/Handler" +
		"\n")
}