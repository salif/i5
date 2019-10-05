// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Block struct {
	Token types.Token
	Body  []Statement
}

func (b Block) Value() string {
	return b.Token.Value
}

func (b Block) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, s := range b.Body {
		out.WriteString(s.String() + ";")
	}
	out.WriteString("}")
	return out.String()
}

func (b Block) statement() {}
