// SPDX-License-Identifier: GPL-3.0-or-later
package object

type String struct {
	Value string
}

func (s *String) Type() TYPE {
	return STRING
}

func (s *String) StringValue() string {
	return s.Value
}
