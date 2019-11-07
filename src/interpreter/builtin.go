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
		return builtin
	} else {
		return newError(true, node.GetLine(), constants.ERROR_REFERENCE, "buitin not found: "+node.GetValue())
	}
}
