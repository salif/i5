// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalBlock(node ast.Block, env *object.Env) object.Object {
	for _, statement := range node.GetBody() {
		var result object.Object = Eval(statement, env)
		if ErrorType(result) == FATAL {
			return result
		}
		if result.Type() == object.RETURN {
			return result
		}
	}
	return Nil(node.GetLine())
}
