// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIf(node ast.If, env *object.Env) error {
	ev, err := Eval(node.GetCondition(), env)

	if err != nil {
		return err
	}

	if ev.Type() != constants.TYPE_BOOLEAN {
		return constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_IS_NOT_A_BOOL, ev.Type())}
	}

	if isTrue(ev) {
		_, err := Eval(node.GetConsequence(), env)
		if err != nil {
			return err
		}
	} else if node.HaveAlternative() {
		_, err := Eval(node.GetAlternative(), env)
		if err != nil {
			return err
		}
	}

	return nil
}
