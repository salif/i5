// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/ast"
)

type Function struct {
	Params []ast.Identifier
	Body   ast.Node
	Env    *Env
}

func (this Function) Type() TYPE {
	return FUNCTION
}

func (this Function) StringValue() string {
	return fmt.Sprintf("[type: %v]", this.Type())
}
