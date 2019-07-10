package parser

import (
	"github.com/i5-lang/i5/src/types"
)

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
	walk(&root.Body, "eof")
	root.Body = append(root.Body, types.Node{
		Kind:  "eol",
		Value: "eol",
	})
	var dn types.Node
	group(root, &dn)
	return dn
}

func walk(cnode *[]types.Node, until string) {
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
		walk(&n.Body, ")")

		*cnode = append(*cnode, n)
	} else if tkn.Kind == "{" {
		n := types.Node{
			Kind: "obj",
			Dlm:  "eol",
			Body: []types.Node{},
		}
		walk(&n.Body, "}")
		*cnode = append(*cnode, n)
	} else {
		*cnode = append(*cnode, types.Node{
			Kind:  tkn.Kind,
			Value: tkn.Value,
		})
	}

	if tkn.Kind != until {
		walk(cnode, until)
	}
}

func group(nd types.Node, dn *types.Node) {
	er := types.Node{
		Kind: "statement",
		Body: []types.Node{},
	}
	if len(nd.Body) == 0 {
		dn.Kind = nd.Kind
		dn.Body = append(dn.Body, er)
	}
	for i, n := range nd.Body {
		if n.Kind == nd.Dlm {
			dn.Kind = nd.Kind
			dn.Body = append(dn.Body, er)
			er = types.Node{
				Kind: "statement",
				Body: []types.Node{},
			}
		} else {
			if n.Dlm == "" {
				er.Body = append(er.Body, n)
			} else {
				var tdn types.Node
				group(n, &tdn)
				er.Body = append(er.Body, tdn)
			}
			if i == len(nd.Body)-1 {
				dn.Kind = nd.Kind
				dn.Body = append(dn.Body, er)
				er = types.Node{
					Kind: "statement",
					Body: []types.Node{},
				}
			}
		}
	}
}
