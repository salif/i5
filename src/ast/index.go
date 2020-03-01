// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"
)

type Index struct {
	line     uint32
	left     Node
	operator string
	right    Node
}

func (this Index) GetType() string {
	return INDEX
}

func (this Index) Debug() string {
	var result strings.Builder
	result.WriteString("(")
	result.WriteString(this.left.Debug())
	result.WriteString(this.operator)
	result.WriteString(this.right.Debug())
	result.WriteString(")")
	return result.String()
}

func (this Index) GetLine() uint32 {
	return this.line
}

func (this Index) GetLeft() Node {
	return this.left
}

func (this Index) GetRight() Node {
	return this.right
}

func (this Index) GetOperator() string {
	return this.operator
}

func (this Index) Init(line uint32, left Node, operator string) Index {
	this.line = line
	this.left = left
	this.operator = operator
	return this
}

func (this *Index) SetRight(right Node) {
	this.right = right
}
