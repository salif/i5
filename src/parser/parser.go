package parser

import "github.com/i5/i5/src/types"

var (
	current int = -1
	tokens  types.TokenList
)

func Run(_tokens types.TokenList) types.Node {
	current = -1
	tokens = _tokens
	var rootNode types.Node = types.Node{
		Kind: "program",
		Dlm:  "eol",
		Body: []types.Node{},
	}
	Walk(&rootNode.Body, "eof")
	rootNode.Body = append(rootNode.Body, types.Node{
		Kind:  "eol",
		Value: "eol",
	})
	var newRootNode types.Node
	Group(rootNode, &newRootNode)
	return newRootNode
}
