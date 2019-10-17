// SPDX-License-Identifier: GPL-3.0-or-later
package constants

const (
	TEST_GOT_WANT string = "error: got: %v, want: %v"

	ARGS_UNKNOWN     string = "unknown option: %v"
	ARGS_UNKNOWN_CLR string = "unknown output format: %v"
	ARGS_NO_FILE     string = "no file specified"

	LEXER_UNEXPECTED_TOKEN string = "line %d: unexpected token '%v'"
	LEXER_OUT_OF_RANGE     string = "line %d: %v: index out of range"

	PARSER_EXPECTED_FOUND string = "line %d: expected '%v', found '%v'"
	PARSER_UNEXPECTED     string = "line %d: unexpected '%v'"
	PARSER_EXPECTED       string = "line %d: expected '%v'"
	PARSER_NOT_INT        string = "line %d: could not parse %v as integer"
	PARSER_NOT_FLOAT      string = "line %d: could not parse %v as float"
	PARSER_EXPECTED_ARG   string = "line %d: expected ',' or ')', found '%v'"

	FILE_NOT_FOUND        string = "%v: no such file or directory"
	FILE_FOUND_FILE       string = "%v: expected directory, found file"
	FILE_FOUND_DIR        string = "%v: expected file, found directory"
	FILE_CANNOT_READ_FILE string = "%v: can not read file"
	FILE_CANNOT_READ_DIR  string = "%v: can not read directory"
	FILE_FILE_EXISTS      string = "%v: file already exists"
	FILE_DIR_EXISTS       string = "%v: directory already exists"
	FILE_CANNOT_WRITE     string = "%v: can not write file"

	IR_INVALID_MOD_FILE  string = "mod file is invalid"
	IR_MAIN_FN_NOT_FOUND string = "main function not found"
	IR_INVALID_EVAL      string = "invalid expression: '%v'"
	IR_INVALID_INFIX     string = "invalid expression: '%v%v%v'"
	IR_INVALID_PREFIX    string = "invalid expression: '%v%v'"
	IR_INVALID_POSTFIX   string = "invalid expression: '%v%v'"
	IR_INVALID_CALL      string = "invalid function caller: '%v'"
	IR_BUILTIN_NOT_FOUND string = "builtin not found: '$%v'"
	IR_NOT_IMPLEMENTED   string = "'%v' not implemented yet"
	IR_CANNOT_ASSIGN     string = "can not assign to %v"
	IR_NON_BOOL          string = "non-bool '%v' used as %v condition"
	IR_NOT_ENOUGH_ARGS   string = "not enough arguments to call function"
	IR_MAP_KEY_NOT_FOUND string = "'%v' not found"
)
