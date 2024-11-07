package netcat

import "strings"

/*
* First replace all unallowed characters
* trim the external spaces to leave a lean message
 */
func trimSpace(s string) string {
	s = replaceSpecialcharacters(s)
	s = strings.TrimSpace(s)
	return s
}

/*
* Replace all special characters with characters that can be recognized and processed by golang
* Make them fit inside a rune
 */
 func replaceSpecialcharacters(s string) string {
	replacer := strings.NewReplacer(
		"\\v", " ",
		"\\n", " ",
		"\\t", " ",
		"\\b", " ",
		"\\r", " ",
		"\\a", " ",
		"\\f", " ",
		"\\x20", " ",
	)
	return replacer.Replace(s)
}
