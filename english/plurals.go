package english

import (
	"strings"
	"unicode"
)

var pluralExceptions = map[string]string{
	"deer":  "deer",
	"moose": "moose",
	"elk":   "elk",

	"sheep":    "sheep",
	"mouse":    "mice",
	"dormouse": "dormice",
	"louse":    "lice",

	"shrimp": "shrimp",
	"squid":  "squid",
	"salmon": "salmon",
	"carp":   "carp",
	"trout":  "trout",

	"ox":    "oxen",
	"goose": "geese",

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

	"criterion":  "criteria",
	"automaton":  "automata",
	"phenomenon": "phenomena",

	"radius":  "radii",
	"fungus":  "fungi",
	"alumnus": "alumni",

	"chassis": "chassis",
	"canto":   "cantos",
	"portico": "porticos",
	"quarto":  "quartos",
	"kimono":  "kimonos",

	"stadium": "stadiums",
}

// pluralRule records the information needed to generate a plural
type pluralRule struct {
	suffix  string // the identifying suffix
	trimLen int    // how many letters to trim before adding the plural
	plural  string // the letters to add to form the plural
}

var pluralRules = []pluralRule{
	{"ff", 0, "s"},
	{"f", 1, "ves"},
	{"fe", 2, "ves"},

	{"ay", 0, "s"},
	{"ey", 0, "s"},
	{"iy", 0, "s"},
	{"oy", 0, "s"},
	{"quy", 1, "ies"},
	{"uy", 0, "s"},
	{"y", 1, "ies"},

	{"is", 2, "es"},
	{"s", 0, "es"},

	{"ix", 1, "ces"},
	{"ex", 2, "ices"},
	{"x", 0, "es"},

	{"z", 0, "es"},

	{"o", 0, "es"},

	{"sh", 0, "es"},
	{"ch", 0, "es"},
}

// plural transforms a word into its plural form
func plural(word string) string {
	// handle exceptions first ...
	plural, ok := pluralExceptions[word]
	if ok {
		return plural
	}

	// ... then the non-standard rules ...
	for _, rule := range pluralRules {
		if strings.HasSuffix(word, rule.suffix) {
			return word[:len(word)-rule.trimLen] + rule.plural
		}
	}

	// ... then the catch-all rule
	return word + "s"
}

// Plural returns the appropriate plural form of the word, depending on the
// value of count. If count is 1 then the word is returned unchanged,
// otherwise the plural form is returned. For instance:
//
//	Plural("error", 1)
//
// would return "error", but
//
//	Plural("error", 2)
//
// would return "errors".
//
// Note that English has lots of rules and edge cases for generating plurals;
// it's not just case of adding an 's' at the end. Additionally there are
// several exceptional words which have their own idiosyncratic plurals. A
// final complication lies with those words whose plural depends on the
// meaning of the word; for instance "medium" meaning a substance through
// which impressions or forces are conveyed has a plural of "media" whereas
// with a meaning of a person who claims to communicate with the dead, the
// plural is "mediums". If you come across a word where you think I have
// missed an exception or there is some general rule which is not followed
// please submit a bug report.
func Plural(word string, count int) string {
	if len(word) == 0 {
		return word
	}

	if count == 1 { // gramatically, any count other than one is plural
		return word
	}

	var initialCapital, allCapitals bool

	initialCapital = unicode.IsUpper([]rune(word)[0])
	if initialCapital {
		if word == strings.ToUpper(word) {
			allCapitals = true
		}
	}

	word = plural(strings.ToLower(word))

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
