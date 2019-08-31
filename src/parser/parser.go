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
		Dlm:  types.EOL,
		Body: []types.Node{},
	}
	Walk(&rootNode.Body, types.EOF, types.EOF)
	rootNode.Body = append(rootNode.Body, types.Node{
		Kind:  types.EOL,
		Value: types.EOL,
	})
	var newRootNode types.Node
	Group(rootNode, &newRootNode)
	return newRootNode
}
