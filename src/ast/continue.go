// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Continue struct {
	Line  int
	Value string
}

func (this Continue) StringValue() string {
	return this.Value
}

func (this Continue) GetLine() int {
	return this.Line
}

func (this Continue) statement() {}
