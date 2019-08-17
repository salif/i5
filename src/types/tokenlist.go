package types

type TokenList struct {
	list []Token
}

func (ta TokenList) Init() {
	ta.list = make([]Token, 0)
}

func (ta *TokenList) Add(kind string, char string, line int) {
	ta.list = append(ta.list, Token{Kind: kind, Value: char, Line: line})
}

func (ta TokenList) Get(index int) Token {
	return ta.list[index]
}

func (ta TokenList) Size() int {
	return len(ta.list)
}
