package src

import (
	"github.com/salifm/iota9/src/file"
	"github.com/salifm/iota9/src/help"
	"github.com/salifm/iota9/src/lexer"
	"github.com/salifm/iota9/src/parser"
	"github.com/salifm/iota9/src/interpreter"
)

func Run(args []string) int {
	return ParseArgs(args);
}

func ParseArgs(args []string) int {
	if len(args) == 1 {
		return help.Print()
	}
	switch args[1] {
	case "--help":
		return help.Print()
	default:
		return Execute(file.Read(args[1]))
	}
}

func Execute(source string) int {
	var tokens []string = lexer.Tokenize(source)
	var ast []string = parser.Parse(tokens)
	return interpreter.Interprete(ast)
}