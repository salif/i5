// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"fmt"

	"github.com/i5/i5/src/ast"
)

func PrintAst(_ast ast.Node) {
	fmt.Print(_ast.Debug())
	fmt.Println()
}
