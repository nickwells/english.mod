package english

import (
	"strings"
	"unicode"
)

var specialCasePlurals = map[string]string{
	"deer":   "deer",
	"sheep":  "sheep",
	"mouse":  "mice",
	"louse":  "lice",
	"moose":  "moose",
	"shrimp": "shrimp",
	"ox":     "oxen",
	"goose":  "geese",

	"man":    "men",
	"woman":  "women",
	"child":  "children",
	"person": "people",

	"tooth": "teeth",
	"foot":  "feet",

	"roof":   "roofs",
	"belief": "beliefs",
	"chef":   "chefs",
	"chief":  "chiefs",
	"cafe":   "cafes",

	"series":  "series",
	"species": "species",

	"die": "dice",

	"photo": "photos",
	"piano": "pianos",
	"halo":  "halos",

	"gas": "gasses",
	"fez": "fezzes",
}

// Plural returns the appropriate plural form of the word, depending on the
// value of n. If n is 1 then the word is returned unchanged, otherwise the
// plural form is returned. For instance:
//
//    Plural("error", 1)
//
// would return "error", but
//
//    Plural("error", 2)
//
// would return "errors"
func Plural(word string, n int) string {
	if len(word) == 0 {
		return word
	}

	var initialCapital bool
	var allCapitals bool
	initialCapital = unicode.IsUpper([]rune(word)[0])
	if initialCapital {
		if word == strings.ToUpper(word) {
			allCapitals = true
		}
	}

	word = plural(strings.ToLower(word), n)
	if initialCapital {
		if allCapitals {
			return strings.ToUpper(word)
		}
		wordRunes := []rune(word)
		wordRunes[0] = unicode.ToUpper(wordRunes[0])
		return string(wordRunes)
	}
	return word
}

type pluralRule struct {
	suffix  string
	trimLen int
	plural  string
}

var pluralRules = []pluralRule{
	{"ff", 0, "s"},
	{"f", 1, "ves"},
	{"fe", 2, "ves"},
	{"is", 2, "es"},
	{"on", 2, "a"},
	{"ay", 0, "s"},
	{"ey", 0, "s"},
	{"iy", 0, "s"},
	{"oy", 0, "s"},
	{"uy", 0, "s"},
	{"y", 1, "ies"},
	{"s", 0, "es"},
	{"x", 0, "es"},
	{"z", 0, "es"},
	{"o", 0, "es"},
	{"sh", 0, "es"},
	{"ch", 0, "es"},
}

func plural(word string, n int) string {
	if n == 1 {
		return word
	} // note for constructing plurals 0 is plural (!)

	// Special cases
	plural, ok := specialCasePlurals[word]
	if ok {
		return plural
	}

	for _, rule := range pluralRules {
		if strings.HasSuffix(word, rule.suffix) {
			return word[:len(word)-rule.trimLen] + rule.plural
		}
	}

	return word + "s"
}
