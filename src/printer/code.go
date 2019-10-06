// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"strings"

	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/types"
)

func Code(tokens types.TokenList) {
	const tab string = "    "
	var tabs int = 0
	var output strings.Builder
	output.WriteString(console.Format("%3d ", 1))

	for i := 0; i < tokens.Size(); i++ {
		var token types.Token = tokens.Get(i)

		if _, isKeyword := lexer.IsKeyword(token.Value); isKeyword {
			output.WriteString(console.Red(token.Value) + " ")
		} else if token.Type == types.IDENTIFIER {
			output.WriteString(console.Green(token.Value))
			if tokens.Get(i+1).Type == types.IDENTIFIER {
				output.WriteString(" ")
			}
		} else if token.Type == types.STRING {
			output.WriteString(console.Yellow("\"" + token.Value + "\""))
		} else if token.Type == types.NUMBER {
			output.WriteString(console.Magenta(token.Value))
		} else if token.Type == types.BUILTIN {
			output.WriteString(console.Cyan("$" + token.Value))
		} else if token.Type == types.EOL {
			output.WriteString("\n")
			output.WriteString(console.Format("%3d ", token.Line+1))
			output.WriteString(strings.Repeat(tab, tabs))
		} else if token.Type == types.EOF {
		} else if token.Type == types.LBRACE {
			output.WriteString(" {")
			tabs++
		} else if token.Type == types.RBRACE {
			output.WriteString("\u0008\u0008\u0008\u0008")
			output.WriteString("} ")
			tabs--
			if tabs < 0 {
				tabs = 0
			}
		} else if token.Type == types.LPAREN || token.Type == types.RPAREN || token.Type == types.DOT {
			output.WriteString(token.Value)
		} else if token.Type == types.COMMA {
			output.WriteString(token.Value + " ")
		} else {
			output.WriteString(console.Red(" " + token.Value + " "))
		}
	}
	output.WriteString("\n")
	console.Println(output.String())
}
