// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/builtins"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalBuiltin(node ast.Builtin, env *object.Env) object.Object {
	if builtin, ok := builtins.Get(node.GetValue(), env); ok {
		if ErrorType(builtin) > 0 {
			builtin := builtin.(object.Error)
			builtin.Line = node.GetLine()
			return builtin
		}
		return builtin
	} else {
		return newError(true, node.GetLine(), constants.ERROR_NIL, "buitin not found: "+node.GetValue())
	}
}
