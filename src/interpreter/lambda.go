// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalLambda(node ast.Lambda, env *object.Env) object.Object {
	return object.Function{Params: node.GetParams(), Body: node.GetBody(), Env: env}
}
