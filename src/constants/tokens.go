// SPDX-License-Identifier: GPL-3.0-or-later
package constants

type Token struct {
	Type  string
	Value string
	Line  uint32
}

const (
	TOKEN_EOL string = "EOL"
	TOKEN_EOF string = "EOF"

	// types
	TOKEN_INTEGER    string = "integer"
	TOKEN_FLOAT      string = "float"
	TOKEN_STRING     string = "string"
	TOKEN_BUILTIN    string = "builtin"
	TOKEN_IDENTIFIER string = "identifier"

	// keywords
	TOKEN_ANDAND string = "and"
	TOKEN_OROR   string = "or"
	TOKEN_FN     string = "fn"
	TOKEN_RETURN string = "return"
	TOKEN_IMPORT string = "import"
	TOKEN_AS     string = "as"
	TOKEN_IF     string = "if"
	TOKEN_ELIF   string = "elif"
	TOKEN_ELSE   string = "else"
	TOKEN_SWITCH string = "switch"
	TOKEN_CASE   string = "case"
	TOKEN_LOOP   string = "loop"
	TOKEN_BREAK  string = "break"
	TOKEN_THROW  string = "throw"

	// assignment operators
	TOKEN_EQ         string = "="
	TOKEN_COLONEQ    string = ":="
	TOKEN_PLUSEQ     string = "+="
	TOKEN_MINUSEQ    string = "-="
	TOKEN_MULTIPLYEQ string = "*="
	TOKEN_DIVIDEEQ   string = "/="
	TOKEN_MODULOEQ   string = "%="
	TOKEN_QM         string = "?"

	// comparison operators
	TOKEN_EQEQ  string = "=="
	TOKEN_NOTEQ string = "!="
	TOKEN_LT    string = "<"
	TOKEN_GT    string = ">"
	TOKEN_LTEQ  string = "<="
	TOKEN_GTEQ  string = ">="

	// arithmetic operators
	TOKEN_PLUS     string = "+"
	TOKEN_MINUS    string = "-"
	TOKEN_MULTIPLY string = "*"
	TOKEN_DIVIDE   string = "/"
	TOKEN_MODULO   string = "%"

	// bitwise operators
	TOKEN_AND  string = "&"
	TOKEN_OR   string = "|"
	TOKEN_XOR  string = "^"
	TOKEN_BNOT string = "~"
	TOKEN_LTLT string = "<<"
	TOKEN_GTGT string = ">>"

	// other operators
	TOKEN_NOT        string = "!"
	TOKEN_COLON      string = ":"
	TOKEN_EQGT       string = "=>"
	TOKEN_QMQM       string = "??"
	TOKEN_COLONCOLON string = "::"
	TOKEN_DOT        string = "."
	TOKEN_COMMA      string = ","

	// brackets
	TOKEN_LPAREN string = "("
	TOKEN_RPAREN string = ")"
	TOKEN_LBRACE string = "{"
	TOKEN_RBRACE string = "}"
)
