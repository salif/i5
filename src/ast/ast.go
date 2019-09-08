// Adapted from https://github.com/prologic/monkey-lang/blob/v1.3.5/ast/ast.go
package ast

type Node interface {
	Value() string
	String() string
}

type Statement interface {
	Node
	statement()
}

type Expression interface {
	Node
	expression()
}
