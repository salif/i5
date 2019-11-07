// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"fmt"
	"strings"

	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/types"
)

func PrintTokens(tokens types.TokenList) {
	var output strings.Builder
	output.WriteString(fmt.Sprintln(" ",
		colors.Cyan("Line"),
		colors.Red("Type"),
		colors.Yellow("Value")))
	output.WriteString("\n")
	for i := 0; i < tokens.Size(); i++ {
		var tkn types.Token = tokens.Get(i)
		output.WriteString(
			fmt.Sprintln(
				colors.Cyan(fmt.Sprintf("%3d ", tkn.Line)),
				pad(colors.Red(tkn.Type)),
				colors.Yellow(tkn.Value)))
	}
	fmt.Println(output.String())
}

func pad(str string) string {
	return str + strings.Repeat(" ", 24-len(str))
}
