package utils

import "strings"

//Sanitize cleans a string
func Sanitize(name string) string {
	sane := strings.ReplaceAll(name, ":", "")
	sane = strings.ReplaceAll(sane, " ", "")
	sane = strings.ReplaceAll(sane, "-", "")
	sane = strings.ReplaceAll(sane, "!", "")
	sane = strings.ReplaceAll(sane, "(", "")
	sane = strings.ReplaceAll(sane, ")", "")
	sane = strings.ReplaceAll(sane, "[", "")
	sane = strings.ReplaceAll(sane, "]", "")
	sane = strings.ReplaceAll(sane, "'", "")
	sane = strings.ReplaceAll(sane, ",", "")
	return sane
}
