// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/builtins"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalBuiltin(node ast.Builtin, env *object.Env) (object.Object, error) {
	ev, ok := builtins.Get(node.GetValue(), env)
	if ok {
		return ev, nil
	} else {
		return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_BUILTIN_NOT_FOUND, node.GetValue())}
	}
}
