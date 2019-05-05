package src

func Run(args []string) int {
	return parseArgs(args)
}

func parseArgs(args []string) int {
	if len(args) == 1 {
		return PrintHelp()
	}
	switch args[1] {
	case "--help":
		return PrintHelp()
	default:
		return Execute(ReadFile(args[1]))
	}
}

func Execute(source string) int {
	var tokens []token = Tokenizer(source)
	var ast []token = Parser(tokens)
	return Interpreter(ast)
}
