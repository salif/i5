// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIdentifier(node ast.Identifier, env *object.Env) object.Object {
	if val, ok := env.Get(node.GetValue()); ok {
		return val
	} else {
		return newError(true, node.GetLine(), constants.ERROR_NIL, "identifier not found: "+node.GetValue())
	}
}
