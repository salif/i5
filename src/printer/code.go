package printer

import (
	"strings"

	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/types"
)

func Code(tokens types.TokenList) {
	const tab string = "    "
	var tabs int = 0
	var output strings.Builder
	output.WriteString(errors.F("%3d ", 1))

	for i := 0; i < tokens.Size(); i++ {
		var token types.Token = tokens.Get(i)

		if _, isKeyword := lexer.IsKeyword(token.Value); isKeyword {
			output.WriteString(console.Color{Value: token.Value}.Red().String() + " ")
		} else if token.Type == types.IDENTIFIER {
			output.WriteString(console.Color{Value: token.Value}.Green().String())
		} else if token.Type == types.STRING {
			output.WriteString(console.Color{Value: "\"" + token.Value + "\""}.Yellow().String())
		} else if token.Type == types.NUMBER {
			output.WriteString(console.Color{Value: token.Value}.Magenta().String())
		} else if token.Type == types.BUILTIN {
			output.WriteString(console.Color{Value: token.Value}.Cyan().String())
		} else if token.Type == types.META {
			output.WriteString(" " + console.Color{Value: token.Value}.Red().String())
		} else if token.Type == types.EOL {
			output.WriteString("\n")
			output.WriteString(errors.F("%3d ", token.Line+1))
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
			output.WriteString(console.Color{Value: " " + token.Value + " "}.Red().String())
		}
	}
	output.WriteString("\n")
	console.Println(output.String())
}
