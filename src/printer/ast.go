package printer

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/io/console"
)

func Ast(_ast ast.Node, tabs int, _color string) {
	console.Println(_ast.String())
}
