package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/io/console"
)

func Run(program ast.Node) {
	switch program.(type) {
	case ast.Program:
		console.Println(console.Color{Value: "interpreter not completed yet"}.Blue())
	}
}
