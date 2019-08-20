package parser

import "github.com/i5/i5/src/types"

func Group(nd types.Node, dn *types.Node) {
	var er types.Node = types.Node{
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
				Group(n, &tdn)
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
