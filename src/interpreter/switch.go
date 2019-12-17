// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalSwitch(node ast.Switch, env *object.Env) error {
	for _, c := range node.GetCases() {
		ev, err := Eval(ast.Infix{}.Set(c.GetLine(), c.GetCase(), constants.TOKEN_EQEQ, node.GetCondition()), env)
		if err != nil {
			return err
		}
		if isTrue(ev) {
			_, errr := Eval(c.GetBody(), env)
			if errr != nil {
				return errr
			}
		}
	}
	return nil
}
