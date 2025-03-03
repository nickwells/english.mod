package english

var ordSfx = [10]string{
	"th", "st", "nd", "rd", "th", "th", "th", "th", "th", "th",
}

// OrdinalSuffix returns the two letter suffix appropriate to the ordinal
// value of n. 1 will give "st", 2 will give "nd" and so on
func OrdinalSuffix(n int) string {
	if n < 0 {
		n *= -1
	}

	nMod100 := n % 100 //nolint:mnd
	switch nMod100 {
	case 11, 12, 13: //nolint:mnd
		return "th"
	}

	nMod10 := n % 10 //nolint:mnd

	return ordSfx[nMod10]
}
