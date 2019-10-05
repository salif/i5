// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"strings"

	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func Tokens(tokens types.TokenList) {
	var output strings.Builder
	output.WriteString(errors.LN(" ",
		console.Color{Value: "Line"}.Cyan(),
		console.Color{Value: "Type"}.Red(),
		console.Color{Value: "Value"}.Yellow()))
	output.WriteString("\n")
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		output.WriteString(errors.LN(console.Color{Value: errors.F("%3d ", tkn.Line)}.Cyan(),
			pad(console.Color{Value: tkn.Type}.Red().String()),
			console.Color{Value: tkn.Value}.Yellow()))
	}
	console.Println(output.String())
}

func pad(str string) string {
	return str + strings.Repeat(" ", 24-len(str))
}
