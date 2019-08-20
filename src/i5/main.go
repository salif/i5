package i5

import (
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
)

func Run(args []string) {
	parseArgs(args)
}

func parseArgs(args []string) {
	if len(args) == 1 {
		PrintHelp()
	}
	switch args[1] {
	case "--help":
		PrintHelp()
	case "--code":
		PrintCode(lexer.Run(file.Read(args[2])))
	case "--tokens":
		PrintTokens(lexer.Run(file.Read(args[2])))
	case "--ast":
		PrintAst(parser.Run(lexer.Run(file.Read(args[2]))))
	default:
		Execute(args[1])
	}
}

func Execute(fileName string) {
	interpreter.Run(parser.Run(lexer.Run(file.Read(fileName))))
}
