package parser

import "github.com/i5/i5/src/types"

func Group(rootNode types.Node, newRootNode *types.Node) {
	var expr types.Node = types.Node{
		Kind: "expr",
		Body: []types.Node{},
	}
	if len(rootNode.Body) == 0 {
		newRootNode.Kind = rootNode.Kind
		newRootNode.Body = append(newRootNode.Body, expr)
	}
	for i, nodeitem := range rootNode.Body {
		if nodeitem.Kind == rootNode.Dlm {
			newRootNode.Kind = rootNode.Kind
			newRootNode.Body = append(newRootNode.Body, expr)
			expr = types.Node{
				Kind: "expr",
				Body: []types.Node{},
			}
		} else {
			if nodeitem.Dlm == "" {
				expr.Body = append(expr.Body, nodeitem)
			} else {
				var tdn types.Node
				Group(nodeitem, &tdn)
				expr.Body = append(expr.Body, tdn)
			}
			if i == len(rootNode.Body)-1 {
				newRootNode.Kind = rootNode.Kind
				newRootNode.Body = append(newRootNode.Body, expr)
				expr = types.Node{
					Kind: "expr",
					Body: []types.Node{},
				}
			}
		}
	}
}
