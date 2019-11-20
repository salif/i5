// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalProgram(node ast.Program, env *object.Env) object.Object {
	for _, fn := range node.GetBody() {
		var result object.Object = Eval(fn, env)
		if ErrorType(result) == FATAL {
			return result
		}
	}
	return Nil(node.GetLine())
}
