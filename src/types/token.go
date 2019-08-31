package types

type Token struct {
	Kind  string
	Value string
	Line  int
}

const (
	BRACKET    = "bracket"
	COMMA      = "comma"
	DOT        = "dot"
	OPERATOR   = "operator"
	NUMBER     = "number"
	STRING     = "string"
	BUILTIN    = "builtin"
	KEYWORD    = "keyword"
	IDENTIFIER = "identifier"
	EOL        = "eol"
	EOF        = "eof"
)
