package english_test

import (
	"testing"

	"github.com/nickwells/english.mod/english"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestJoin(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		s   []string
		exp string
	}{
		{
			ID: testhelper.MkID("empty slice"),
		},
		{
			ID:  testhelper.MkID("one entry"),
			s:   []string{"one"},
			exp: "one",
		},
		{
			ID:  testhelper.MkID("two entries"),
			s:   []string{"one", "two"},
			exp: "one or two",
		},
		{
			ID:  testhelper.MkID("three entries"),
			s:   []string{"one", "two", "three"},
			exp: "one, two or three",
		},
	}

	for _, tc := range testCases {
		result := english.Join(tc.s, ", ", " or ")
		if result != tc.exp {
			t.Log(tc.IDStr())
			t.Log("\t: expected: ", tc.exp)
			t.Log("\t:      got: ", result)
			t.Errorf("\t: Unexpected result\n")
		}
	}
}

func TestJoinQuoted(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		s      []string
		quotes []string
		exp    string
	}{
		{
			ID:     testhelper.MkID("empty slice. Two, differing quotes"),
			quotes: []string{"\u00ab", "\u00bb"},
		},
		{
			ID:     testhelper.MkID("one entry. Two, differing quotes"),
			s:      []string{"one"},
			quotes: []string{"\u00ab", "\u00bb"},
			exp:    "«one»",
		},
		{
			ID:     testhelper.MkID("two entries. Two, differing quotes"),
			s:      []string{"one", "two"},
			quotes: []string{"\u00ab", "\u00bb"},
			exp:    "«one» or «two»",
		},
		{
			ID:     testhelper.MkID("three entries. Two, differing quotes"),
			s:      []string{"one", "two", "three"},
			quotes: []string{"\u00ab", "\u00bb"},
			exp:    "«one», «two» or «three»",
		},
		{
			ID:     testhelper.MkID("three entries. Three, differing quotes"),
			s:      []string{"one", "two", "three"},
			quotes: []string{"\u00ab", "\u00bb", "<>"},
			exp:    "«one», «two» or «three»",
		},
		{
			ID:     testhelper.MkID("empty slice. One quote value"),
			quotes: []string{"'"},
		},
		{
			ID:     testhelper.MkID("one entry. One quote value"),
			s:      []string{"one"},
			quotes: []string{"'"},
			exp:    "'one'",
		},
		{
			ID:     testhelper.MkID("two entries. One quote value"),
			s:      []string{"one", "two"},
			quotes: []string{"'"},
			exp:    "'one' or 'two'",
		},
		{
			ID:     testhelper.MkID("three entries. One quote value"),
			s:      []string{"one", "two", "three"},
			quotes: []string{"'"},
			exp:    "'one', 'two' or 'three'",
		},
		{
			ID: testhelper.MkID("empty slice. No quote values"),
		},
		{
			ID:  testhelper.MkID("one entry. No quote values"),
			s:   []string{"one"},
			exp: `"one"`,
		},
		{
			ID:  testhelper.MkID("two entries. No quote values"),
			s:   []string{"one", "two"},
			exp: `"one" or "two"`,
		},
		{
			ID:  testhelper.MkID("three entries. No quote values"),
			s:   []string{"one", "two", "three"},
			exp: `"one", "two" or "three"`,
		},
	}

	for _, tc := range testCases {
		result := english.JoinQuoted(tc.s, ", ", " or ", tc.quotes...)
		if result != tc.exp {
			t.Log(tc.IDStr())
			t.Log("\t: expected: ", tc.exp)
			t.Log("\t:      got: ", result)
			t.Errorf("\t: Unexpected result\n")
		}
	}
}
