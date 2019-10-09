// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
	"github.com/i5/i5/src/types"
)

// Print Code
func Print(name string, isName bool, _code bool, _tokens bool, _ast bool) {
	var tokenList types.TokenList

	if isName {
		tokenList = lexer.Run(file.Read(name))
	} else {
		tokenList = lexer.Run([]byte(name))
	}
	if _code {
		printCode(tokenList)
	}
	if _tokens {
		printTokens(tokenList)
	}
	if _ast {
		printAst(parser.Run(tokenList))
	}
}
