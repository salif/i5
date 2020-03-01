// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package constants

type Token struct {
	Type  string
	Value string
	Line  uint32
}

const (
	// end of line
	TOKEN_EOL string = "EOL"
	// end of file
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

var KEYWORDS = map[string]string{
	TOKEN_ANDAND: TOKEN_ANDAND,
	TOKEN_OROR:   TOKEN_OROR,
	TOKEN_FN:     TOKEN_FN,
	TOKEN_RETURN: TOKEN_RETURN,
	TOKEN_IMPORT: TOKEN_IMPORT,
	TOKEN_AS:     TOKEN_AS,
	TOKEN_IF:     TOKEN_IF,
	TOKEN_ELIF:   TOKEN_IF,
	TOKEN_ELSE:   TOKEN_ELSE,
	TOKEN_SWITCH: TOKEN_SWITCH,
	TOKEN_CASE:   TOKEN_CASE,
	TOKEN_LOOP:   TOKEN_LOOP,
	TOKEN_BREAK:  TOKEN_BREAK,
	TOKEN_THROW:  TOKEN_THROW,
}
