package parser

import "github.com/i5/i5/src/types"

func Walk(node *[]types.Node, untilKind string, untilValue string) {
	current++
	var tkn types.Token = tokens.Get(current)
	if tkn.Kind == untilKind && tkn.Value == untilValue {
		return
	}
	if tkn.Kind == types.BRACKET && tkn.Value == "(" {
		var n types.Node = types.Node{
			Kind: "args",
			Dlm:  "Dlm",
			Body: []types.Node{},
		}
		Walk(&n.Body, types.BRACKET, ")")

		*node = append(*node, n)
	} else if tkn.Kind == types.BRACKET && tkn.Value == "{" {
		var n types.Node = types.Node{
			Kind: "body",
			Dlm:  "eol",
			Body: []types.Node{},
		}
		Walk(&n.Body, types.BRACKET, "}")
		*node = append(*node, n)
	} else {
		*node = append(*node, types.Node{
			Kind:  tkn.Kind,
			Value: tkn.Value,
		})
	}

	if !(tkn.Kind == untilKind && tkn.Value == untilValue) {
		Walk(node, untilKind, untilValue)
	}
}
