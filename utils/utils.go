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
	sane = strings.ReplaceAll(sane, "/", "")
	return sane
}

//Contains string
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
