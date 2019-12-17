// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

type Function struct {
	Params []ast.Identifier
	Body   ast.Node
	Env    *Env
}

func (this Function) Type() string {
	return constants.TYPE_FUNCTION
}

func (this Function) StringValue() string {
	return fmt.Sprintf("[type: %v, params: %v]", this.Type(), len(this.Params))
}
