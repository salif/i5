package parser

import "github.com/i5/i5/src/types"

func Walk(node *[]types.Node, until string) {
	current++
	var tkn types.Token = tokens.Get(current)
	if tkn.Kind == until {
		return
	}
	if tkn.Kind == "(" {
		var n types.Node = types.Node{
			Kind: "args",
			Dlm:  "Dlm",
			Body: []types.Node{},
		}
		Walk(&n.Body, ")")

		*node = append(*node, n)
	} else if tkn.Kind == "{" {
		var n types.Node = types.Node{
			Kind: "body",
			Dlm:  "eol",
			Body: []types.Node{},
		}
		Walk(&n.Body, "}")
		*node = append(*node, n)
	} else {
		*node = append(*node, types.Node{
			Kind:  tkn.Kind,
			Value: tkn.Value,
		})
	}

	if tkn.Kind != until {
		Walk(node, until)
	}
}
