package i5

import (
	"strings"

	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func PrintCode(tokens types.TokenList) {
	const tab string = "    "
	var tabs int = 0
	var output strings.Builder
	output.WriteString(errors.F("%3d ", 1))

	for i := 0; i < tokens.Size(); i++ {
		var token types.Token = tokens.Get(i)
		if token.Kind == types.KEYWORD {
			output.WriteString(console.Color{Value: token.Value}.Red().String() + " ")
		} else if token.Kind == types.IDENTIFIER {
			output.WriteString(console.Color{Value: token.Value}.Green().String())
		} else if token.Kind == types.STRING {
			output.WriteString(console.Color{Value: "\"" + token.Value + "\""}.Yellow().String())
		} else if token.Kind == types.NUMBER {
			output.WriteString(console.Color{Value: token.Value}.Magenta().String())
		} else if token.Kind == types.BUILTIN {
			output.WriteString(console.Color{Value: token.Value}.Cyan().String())
		} else if token.Kind == types.EOL {
			output.WriteString("\n")
			output.WriteString(errors.F("%3d ", token.Line+1))
			output.WriteString(strings.Repeat(tab, tabs))
		} else if token.Kind == types.EOF {
		} else if token.Kind == types.OPERATOR {
			if token.Value == "{" {
				output.WriteString(" {")
				tabs++
			} else if token.Value == "}" {
				output.WriteString("\u0008\u0008\u0008\u0008")
				output.WriteString("} ")
				tabs--
				if tabs < 0 {
					tabs = 0
				}
			} else if token.Value == "(" || token.Value == ")" {
				output.WriteString(token.Value)
			} else {
				output.WriteString(console.Color{Value: " " + token.Value + " "}.Red().String())
			}
		} else {
			output.WriteString(token.Value)
		}
	}
	output.WriteString("\n")
	console.Println(output.String())
}

func PrintTokens(tokens types.TokenList) {
	var output strings.Builder
	output.WriteString(errors.LN(" ",
		console.Color{Value: "Line"}.Cyan(),
		console.Color{Value: "Type"}.Red(),
		console.Color{Value: "Value"}.Yellow()))
	output.WriteString("\n")
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		output.WriteString(errors.LN(console.Color{Value: errors.F("%3d ", tkn.Line)}.Cyan(),
			console.Color{Value: tkn.Kind}.Red(),
			console.Color{Value: tkn.Value}.Yellow()))
	}
	console.Println(output.String())
}

func PrintAst(ast types.Node, tabs int, _color string) {
	const tab string = "    "
	if ast.Value == "" {
		console.Print(strings.Repeat(tab, tabs), console.Color{Value: ast.Kind}.Green())
		if len(ast.Body) == 0 {
			console.Println(console.Color{Value: " {}"}.ValueOf(_color))
		} else {
			console.Println(console.Color{Value: " {"}.ValueOf(_color))
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
			console.Println(console.Color{Value: "}"}.ValueOf(_color))
		}
	} else {
		console.Print(strings.Repeat(tab, tabs))
		console.Println(console.Color{Value: ast.Kind}.Cyan(), ast.Value)
	}
}
