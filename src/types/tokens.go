// SPDX-License-Identifier: GPL-3.0-or-later
package types

const (
	// types
	INT     string = "integer"
	FLOAT   string = "float"
	STRING  string = "string"
	BUILTIN string = "builtin"
	IDENT   string = "identifier"
	EOL     string = "new line"
	EOF     string = "end of file"

	// keywords
	ANDAND string = "and"
	OROR   string = "or"
	IF     string = "if"
	ELIF   string = "elif"
	ELSE   string = "else"
	SWITCH string = "switch"
	CASE   string = "case"
	LOOP   string = "loop"
	RETURN string = "return"
	THROW  string = "throw"

	// assignment operators
	EQ         string = "="
	COLONEQ    string = ":="
	PLUSEQ     string = "+="
	MINUSEQ    string = "-="
	MULTIPLYEQ string = "*="
	DIVIDEEQ   string = "/="
	MODULOEQ   string = "%="
	QM         string = "?"

	// comparison operators
	EQEQ  string = "=="
	NOTEQ string = "!="
	LT    string = "<"
	GT    string = ">"
	LTEQ  string = "<="
	GTEQ  string = ">="

	// arithmetic operators
	PLUS     string = "+"
	MINUS    string = "-"
	MULTIPLY string = "*"
	DIVIDE   string = "/"
	MODULO   string = "%"

	// bitwise operators
	AND  string = "&"
	OR   string = "|"
	XOR  string = "^"
	BNOT string = "~"
	LTLT string = "<<"
	GTGT string = ">>"

	// other operators
	NOT        string = "!"
	COLON      string = ":"
	EQGT       string = "=>"
	QMQM       string = "??"
	COLONCOLON string = "::"
	DOT        string = "."
	COMMA      string = ","

	// brackets
	LPAREN string = "("
	RPAREN string = ")"
	LBRACE string = "{"
	RBRACE string = "}"
)
