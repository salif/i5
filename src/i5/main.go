package i5

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
	"strconv"
	"strings"
)

func Run(args []string) {
	ParseArgs(args)
}

func PrintCode(tokens types.TokenList) {
	const tab string = "    "
	var tabs int = 0
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		var tknk string = tkn.Kind
		if tknk == "keyword" || tknk == "bool" {
			console.Print(console.Color(tkn.Value, "red"), " ")
		} else if tknk == "identifier" {
			console.Print(console.Color(tkn.Value, "green"))
		} else if tknk == "string" {
			console.Print(console.Color("\""+tkn.Value+"\"", "yellow"))
		} else if tknk == "number" {
			console.Print(console.Color(tkn.Value, "magenta"))
		} else if tknk == "builtin" {
			console.Print(console.Color(tkn.Value, "cian"))
		} else if tknk == "operator" {
			console.Print(console.Color(" "+tkn.Value+" ", "red"))
		} else if tknk == "eol" || tknk == "eof" {
			console.Println()
			console.Print(strings.Repeat(tab, tabs))
		} else if tknk == "{" {
			console.Print("{")
			tabs++
		} else if tknk == "}" {
			console.Print("\u0008\u0008\u0008\u0008")
			console.Print("} ")
			tabs--
		} else if tknk == ")" {
			console.Print(") ")
		} else {
			console.Print(tkn.Value)
		}
	}
}

func PrintTokens(tokens types.TokenList) {
	console.Println(console.Color("Line", "cian"), console.Color("Type", "red"), console.Color("Value", "yellow"))
	console.Println()
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		console.Println(
			console.Color(strconv.Itoa(tkn.Line), "cian"), console.Color(tkn.Kind, "red"), console.Color(tkn.Value, "yellow"))
	}
}

func PrintAst(ast types.Node) {
	// TODO
}

func PrintHelp() {
	console.Println(`
Usage:

     i5 [options] [file]

options:

     --help           print help
     --code           print code
     --tokens         print tokens
     --ast            print ast
     `)
	errors.Exit(0)
}
