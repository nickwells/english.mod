package english

import "strings"

// Join joins the strings in s using the sep for all but the final join for
// which it uses finalSep. If s is empty it will return the empty string, if
// it has only one entry it wil return just that entry. This allows you to
// produce joined strings like: "a, b, c or d".
func Join(s []string, sep, finalSep string) string {
	joined := ""

	switch len(s) {
	case 0:
	case 1:
		joined = s[0]
	default:
		joined = strings.Join(s[:len(s)-1], sep)
		joined += finalSep + s[len(s)-1]
	}

	return joined
}

// JoinQuoted joins the strings in s using the sep for all but the final join
// for which it uses finalSep. If s is empty it will return the empty string,
// if it has only one entry it wil return just that entry. Each string in s
// is quoted with the supplied opening and closing quotes. This allows you to
// produce joined strings like: "'a', 'b', 'c' or 'd'".
func JoinQuoted(s []string, sep, finalSep, openQ, closeQ string) string {
	joined := ""

	switch len(s) {
	case 0:
	case 1:
		joined = openQ + s[0] + closeQ
	default:
		joined = openQ + strings.Join(s[:len(s)-1], closeQ+sep+openQ)
		joined += closeQ + finalSep + openQ + s[len(s)-1] + closeQ
	}

	return joined
}
