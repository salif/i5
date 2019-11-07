// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"strings"
)

type Identifiers struct {
	line uint32
	body []Identifier
}

func (this Identifiers) GetType() string {
	return IDENTIFIERS
}

func (this Identifiers) Debug() string {
	var result strings.Builder
	var n []string
	for _, p := range this.body {
		n = append(n, p.Debug())
	}
	result.WriteString(strings.Join(n, " "))
	return result.String()
}

func (this Identifiers) GetLine() uint32 {
	return this.line
}

func (this Identifiers) Init(line uint32) Identifiers {
	this.line = line
	this.body = []Identifier{}
	return this
}

func (this Identifiers) GetBody() []Identifier {
	return this.body
}

func (this *Identifiers) Append(node Identifier) {
	this.body = append(this.body, node)
}
