package i5

import (
	"fmt"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
	"strings"
)

func PrintCode(tokens types.TokenList) {
	const tab string = "    "
	var tabs int = 0
	var output strings.Builder
	output.WriteString(fmt.Sprintf("%3d ", 1))

	for i := 0; i < tokens.Size(); i++ {
		var token types.Token = tokens.Get(i)
		var tkn string = token.Kind
		if tkn == "keyword" || tkn == "bool" {
			output.WriteString(fmt.Sprint(console.Color(token.Value, "red"), " "))
		} else if tkn == "identifier" {
			output.WriteString(fmt.Sprint(console.Color(token.Value, "green")))
		} else if tkn == "string" {
			output.WriteString(fmt.Sprint(console.Color("\""+token.Value+"\"", "yellow")))
		} else if tkn == "number" {
			output.WriteString(fmt.Sprint(console.Color(token.Value, "magenta")))
		} else if tkn == "builtin" {
			output.WriteString(fmt.Sprint(console.Color(token.Value, "cyan")))
		} else if tkn == "operator" {
			output.WriteString(fmt.Sprint(console.Color(" "+token.Value+" ", "red")))
		} else if tkn == "dlm" {
			output.WriteString(fmt.Sprint(", "))
		} else if tkn == "eol" {
			output.WriteString(fmt.Sprintln())
			output.WriteString(fmt.Sprintf("%3d ", token.Line+1))
			output.WriteString(fmt.Sprint(strings.Repeat(tab, tabs)))
		} else if tkn == "eof" {
		} else if tkn == "{" {
			output.WriteString(fmt.Sprint(" {"))
			tabs++
		} else if tkn == "}" {
			output.WriteString(fmt.Sprint("\u0008\u0008\u0008\u0008"))
			output.WriteString(fmt.Sprint("} "))
			tabs--
			if tabs < 0 {
				tabs = 0
			}
		} else if tkn == ")" {
			output.WriteString(fmt.Sprint(")"))
		} else {
			output.WriteString(fmt.Sprint(token.Value))
		}
	}
	output.WriteString(fmt.Sprintln())
	console.Println(output.String())
}

func PrintTokens(tokens types.TokenList) {
	var output strings.Builder
	output.WriteString(fmt.Sprintln(" ", console.Color("Line", "cyan"), console.Color("Type", "red"), console.Color("Value", "yellow")))
	output.WriteString(fmt.Sprintln())
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		output.WriteString(fmt.Sprintln(
			console.Color(fmt.Sprintf("%3d ", tkn.Line), "cyan"), console.Color(tkn.Kind, "red"), console.Color(tkn.Value, "yellow")))
	}
	console.Println(output.String())
}

func PrintAst(ast types.Node, tabs int, _color string) {
	const tab string = "    "
	if ast.Value == "" {
		console.Print(strings.Repeat(tab, tabs), console.Color(ast.Kind, "green"))
		if len(ast.Body) == 0 {
			console.Println(console.Color(" {}", _color))
		} else {
			console.Println(console.Color(" {", _color))
			var _ncolor string
			if _color == "red" {
				_ncolor = "yellow"
			} else {
				_ncolor = "red"
			}
			for i := 0; i < len(ast.Body); i++ {
				PrintAst(ast.Body[i], tabs+1, _ncolor)
			}
			console.Print(strings.Repeat(tab, tabs))
			console.Println(console.Color("}", _color))
		}
	} else {
		console.Print(strings.Repeat(tab, tabs))
		console.Println(console.Color(ast.Kind, "cyan"), ast.Value)
	}
}
