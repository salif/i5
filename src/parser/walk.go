package parser

import "github.com/i5/i5/src/types"

func Walk(cnode *[]types.Node, until string) {
	current++
	tkn := tokens.Get(current)
	if tkn.Kind == until {
		return
	}
	if tkn.Kind == "(" {
		n := types.Node{
			Kind: "call",
			Dlm:  "Dlm",
			Body: []types.Node{},
		}
		Walk(&n.Body, ")")

		*cnode = append(*cnode, n)
	} else if tkn.Kind == "{" {
		n := types.Node{
			Kind: "obj",
			Dlm:  "eol",
			Body: []types.Node{},
		}
		Walk(&n.Body, "}")
		*cnode = append(*cnode, n)
	} else {
		*cnode = append(*cnode, types.Node{
			Kind:  tkn.Kind,
			Value: tkn.Value,
		})
	}

	if tkn.Kind != until {
		Walk(cnode, until)
	}
}
