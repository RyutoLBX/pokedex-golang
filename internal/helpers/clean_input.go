package helpers

import "strings"

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
