package english_test

import (
	"fmt"
	"testing"

	"github.com/nickwells/english.mod/english"
	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestPlurals(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		word    string
		n       int
		expRslt string
	}{
		{
			ID:      testhelper.MkID("empty string"),
			word:    "",
			n:       99,
			expRslt: "",
		},
		{
			ID:      testhelper.MkID("special case"),
			word:    "man",
			n:       0,
			expRslt: "men",
		},
		{
			ID:      testhelper.MkID("special case"),
			word:    "man",
			n:       1,
			expRslt: "man",
		},
		{
			ID:      testhelper.MkID("special case"),
			word:    "man",
			n:       2,
			expRslt: "men",
		},
		{
			ID:      testhelper.MkID("standard"),
			word:    "dog",
			n:       1,
			expRslt: "dog",
		},
		{
			ID:      testhelper.MkID("standard"),
			word:    "dog",
			n:       2,
			expRslt: "dogs",
		},
		{
			ID:      testhelper.MkID("standard - initial cap"),
			word:    "Dog",
			n:       2,
			expRslt: "Dogs",
		},
		{
			ID:      testhelper.MkID("standard - all caps"),
			word:    "DOG",
			n:       2,
			expRslt: "DOGS",
		},
		{
			ID:      testhelper.MkID("f-suffixes - single f"),
			word:    "wolf",
			n:       2,
			expRslt: "wolves",
		},
		{
			ID:      testhelper.MkID("f-suffixes - double f"),
			word:    "cliff",
			n:       2,
			expRslt: "cliffs",
		},
		{
			ID:      testhelper.MkID("f-suffixes - fe"),
			word:    "wife",
			n:       2,
			expRslt: "wives",
		},
		{
			ID:      testhelper.MkID("is-suffixes"),
			word:    "analysis",
			n:       2,
			expRslt: "analyses",
		},
		{
			ID:      testhelper.MkID("on-suffixes"),
			word:    "criterion",
			n:       2,
			expRslt: "criteria",
		},
		{
			ID:      testhelper.MkID("y-suffixes - preceding vowel"),
			word:    "boy",
			n:       2,
			expRslt: "boys",
		},
		{
			ID:      testhelper.MkID("y-suffixes - preceding consonant"),
			word:    "fly",
			n:       2,
			expRslt: "flies",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - s"),
			word:    "bus",
			n:       2,
			expRslt: "buses",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - ss"),
			word:    "cross",
			n:       2,
			expRslt: "crosses",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - sh"),
			word:    "bush",
			n:       2,
			expRslt: "bushes",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - ch"),
			word:    "church",
			n:       2,
			expRslt: "churches",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - x"),
			word:    "fox",
			n:       2,
			expRslt: "foxes",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - x - exception"),
			word:    "ox",
			n:       2,
			expRslt: "oxen",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - z"),
			word:    "klutz",
			n:       2,
			expRslt: "klutzes",
		},
		{
			ID:      testhelper.MkID("misc-suffixes - z - exception"),
			word:    "fez",
			n:       2,
			expRslt: "fezzes",
		},
	}

	for _, tc := range testCases {
		rslt := english.Plural(tc.word, tc.n)
		testhelper.DiffString(t,
			tc.IDStr()+fmt.Sprintf(": %d of: %q", tc.n, tc.word), "plural",
			rslt, tc.expRslt)
	}
}
