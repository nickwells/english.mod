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
// is quoted with the supplied opening and closing quotes. If no quotes are
// given both quotes default to '"'. If only one value is given then both the
// opening and closing quotes are set to this value. If more than one value
// is given the opening quote is set to the first value and the closing quote
// is set to the second; subsequent values are silently ignored. This allows
// you to produce joined strings like: "'a', 'b', 'c' or 'd'".
func JoinQuoted(s []string, sep, finalSep string, quotes ...string) string {
	joined := ""
	openQ := `"`
	closeQ := `"`

	if len(quotes) > 0 {
		openQ = quotes[0]
		closeQ = quotes[0]

		if len(quotes) > 1 {
			closeQ = quotes[1]
		}
	}

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
