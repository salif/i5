// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"fmt"
	"strings"

	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func Tokens(tokens types.TokenList) {
	var output strings.Builder
	output.WriteString(fmt.Sprintln(" ",
		console.Cyan("Line"),
		console.Red("Type"),
		console.Yellow("Value")))
	output.WriteString("\n")
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		output.WriteString(
			fmt.Sprintln(
				console.Cyan(console.Format("%3d ", tkn.Line)),
				pad(console.Red(tkn.Type)),
				console.Yellow(tkn.Value)))
	}
	console.Println(output.String())
}

func pad(str string) string {
	return str + strings.Repeat(" ", 24-len(str))
}
