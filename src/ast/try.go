// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Try struct {
	Token types.Token
	Body  *Block
}

func (t Try) Value() string {
	return t.Token.Value
}

func (t Try) String() string {
	var out bytes.Buffer
	out.WriteString("try ")
	out.WriteString(t.Body.String())
	return out.String()
}
func (t Try) statement() {}
