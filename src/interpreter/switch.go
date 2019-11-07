// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalSwitch(node ast.Switch, env *object.Env) object.Object {
	return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_NOT_IMPLEMENTED, "switch")
	// TODO
}
