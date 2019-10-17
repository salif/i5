// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
)

func ParseModuleFile(f []byte) []string {
	result := []string{}
	var scanner Scanner
	scanner.Init(f, func(length int, position int, line int) {
		console.ThrowSyntaxError(1, constants.LEXER_OUT_OF_RANGE, line, "")
	})

	for scanner.HasNext() {

		if scanner.PeekEquals(10) || scanner.PeekEquals(13) {
			scanner.Next()
			continue
		}

		if scanner.PeekEquals(35) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(10); scanner.Next() {
			}
			continue
		}

		var value string = ""
		for ; scanner.HasNext() && scanner.Until(10); scanner.Next() {
			value += string(scanner.Peek())
		}
		result = append(result, value)
		continue
	}

	return result
}
