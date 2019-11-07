// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"fmt"
	"strings"

	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/types"
)

func PrintCode(tokens types.TokenList) {
	const tab string = "    "
	var tabs int = 0
	var output strings.Builder
	output.WriteString(fmt.Sprintf("%3d ", 1))

	for i := 0; i < tokens.Size(); i++ {
		var token types.Token = tokens.Get(i)

		if _, isKeyword := lexer.IsKeyword(token.Value); isKeyword {
			output.WriteString(colors.Red(token.Value) + " ")
		} else if token.Type == types.IDENT {
			output.WriteString(colors.Green(token.Value))
			if tokens.Get(i+1).Type == types.IDENT {
				output.WriteString(" ")
			}
		} else if token.Type == types.STRING {
			output.WriteString(colors.Yellow("\"" + token.Value + "\""))
		} else if token.Type == types.INT || token.Type == types.FLOAT {
			output.WriteString(colors.Magenta(token.Value))
		} else if token.Type == types.BUILTIN {
			output.WriteString(colors.Cyan("$" + token.Value))
		} else if token.Type == types.EOL {
			output.WriteString("\n")
			output.WriteString(fmt.Sprintf("%3d ", token.Line+1))
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
			output.WriteString(colors.Red(" " + token.Value + " "))
		}
	}
	output.WriteString("\n")
	fmt.Println(output.String())
}
