// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalBreak(node ast.Break, env *object.Env) error {
	return constants.Error{Line: node.GetLine(), Type: constants.ERROR_BREAK, Message: ""}
}
