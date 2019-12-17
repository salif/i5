// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalReturn(node ast.Return, env *object.Env) error {
	ev, err := Eval(node.GetBody(), env)
	if err != nil {
		return err
	}
	return constants.Error{Line: node.GetLine(), Type: constants.ERROR_RETURN, Message: "", Value: ev}
}
