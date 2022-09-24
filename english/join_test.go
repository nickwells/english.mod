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
