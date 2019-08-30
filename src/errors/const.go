package errors

const (
	ARGS_UNKNOWN           string = "unknown option: %v"
	ARGS_UNKNOWN_CLR       string = "unknown colorizer: %v"
	ARGS_NO_FILE           string = "no file specified"
	LEXER_UNEXPECTED_TOKEN string = "line %v: %v: unexpected token\n"
	SCANNER_OUT_OF_RANGE   string = "line %v: %v: index out of range\n"
	READER_NOT_FOUND       string = "%v: no such file or directory"
	TEST_GOT_WANT          string = "error: got: %v, want: %v"
)
