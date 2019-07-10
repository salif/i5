package i5

import (
	"github.com/i5-lang/i5/src/interpreter"
	"github.com/i5-lang/i5/src/lexer"
	"github.com/i5-lang/i5/src/parser"
)

func Run(args []string) {
	parseArgs(args)
}

func parseArgs(args []string) {
	if len(args) == 1 {
		PrintHelp()
		return
	}
	switch args[1] {
	case "--help":
		PrintHelp()
	default:
		execute(args[1])
	}
}

func execute(fileName string) {
	interpreter.Run(parser.Run(lexer.Run(ReadFile(fileName))))
}
