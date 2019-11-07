// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/file"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
)

func Print(name string, tokens bool, code bool, ast bool) error {
	var result int = file.Info(name)
	switch result {
	case 1:
		return fmt.Errorf("%v%v\n", colors.Red("error: "), fmt.Sprintf(constants.FILE_NOT_FOUND, name))
	case 3:
		codeBytes, err := file.Read(name)
		if err != nil {
			return fmt.Errorf("%v%v", colors.Red("error: "), err)
		}
		t, err := lexer.Run(name, codeBytes)
		if err != nil {
			return err
		}
		if code {
			PrintCode(t)
		}
		if tokens {
			PrintTokens(t)
		}
		if ast {
			a, err := parser.Run(name, t)
			if err != nil {
				return err
			}
			PrintAst(a)
		}

		return nil
	default:
		return nil
	}
}
