package console

func Color(text string, color string) string {
	if color == "red" {
		return "\x1b[91m" + text + "\x1b[0m"
	} else if color == "green" {
		return "\x1b[92m" + text + "\x1b[0m"
	} else if color == "yellow" {
		return "\x1b[93m" + text + "\x1b[0m"
	} else if color == "blue" {
		return "\x1b[94m" + text + "\x1b[0m"
	} else if color == "magenta" {
		return "\x1b[95m" + text + "\x1b[0m"
	} else if color == "cian" {
		return "\x1b[96m" + text + "\x1b[0m"
	} else {
		return text
	}
}
