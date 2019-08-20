package parser

import "github.com/i5/i5/src/types"

var (
	current int = -1
	tokens  types.TokenList
)

func Run(tkns types.TokenList) types.Node {
	tokens = tkns
	root := types.Node{
		Kind: "program",
		Dlm:  "eol",
		Body: []types.Node{},
	}
	Walk(&root.Body, "eof")
	root.Body = append(root.Body, types.Node{
		Kind:  "eol",
		Value: "eol",
	})
	var dn types.Node
	Group(root, &dn)
	return dn
}
