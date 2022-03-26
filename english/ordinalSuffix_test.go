package english_test

import (
	"testing"

	"github.com/nickwells/english.mod/english"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestOrdinalSuffix(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		n       int
		expRslt string
	}{
		{
			ID:      testhelper.MkID(""),
			n:       0,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       -1,
			expRslt: "st",
		},
		{
			ID:      testhelper.MkID(""),
			n:       1,
			expRslt: "st",
		},
		{
			ID:      testhelper.MkID(""),
			n:       21,
			expRslt: "st",
		},
		{
			ID:      testhelper.MkID(""),
			n:       201,
			expRslt: "st",
		},
		{
			ID:      testhelper.MkID(""),
			n:       2,
			expRslt: "nd",
		},
		{
			ID:      testhelper.MkID(""),
			n:       3,
			expRslt: "rd",
		},
		{
			ID:      testhelper.MkID(""),
			n:       4,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       5,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       6,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       7,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       8,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       9,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       11,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       111,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       12,
			expRslt: "th",
		},
		{
			ID:      testhelper.MkID(""),
			n:       13,
			expRslt: "th",
		},
	}

	for _, tc := range testCases {
		rslt := english.OrdinalSuffix(tc.n)
		if rslt != tc.expRslt {
			t.Log(tc.IDStr())
			t.Logf("\t: english.OrdinalSuffix(n = %d)\n",
				tc.n)
			t.Logf("\t: expected: %q\n", tc.expRslt)
			t.Logf("\t:      got: %q\n", rslt)
			t.Errorf("\t: Unexpected result\n")
		}
	}
}
