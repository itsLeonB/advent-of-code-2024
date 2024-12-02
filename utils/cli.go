package utils

func CleanInput(input string) string {
	return input[:len(input)-1] // Trim the trailing newline
}
