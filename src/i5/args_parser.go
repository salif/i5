package i5

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
)

func ParseArgs(args []string) {
	if len(args) == 1 {
		PrintHelp()
	}
	switch args[1] {
	case "--help":
		PrintHelp()
	case "--code":
		if len(args) == 2 {
			errors.NewFatalError(errors.ARGS_NO_FILE, 1)
		}
		PrintCode(lexer.Run(file.Read(args[2])))
	case "--tokens":
		if len(args) == 2 {
			errors.NewFatalError(errors.ARGS_NO_FILE, 1)
		}
		PrintTokens(lexer.Run(file.Read(args[2])))
	case "--ast":
		if len(args) == 2 {
			errors.NewFatalError(errors.ARGS_NO_FILE, 1)
		}
		PrintAst(parser.Run(lexer.Run(file.Read(args[2]))))
	default:
		Execute(args[1])
	}
}

func Execute(fileName string) {
	interpreter.Run(parser.Run(lexer.Run(file.Read(fileName))))
}
