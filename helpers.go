package helpers

// TextFromNumber Returns the text based form of a number.
// e.g. 1 would return One
func TextFromNumber(num int) string {
	switch num {
	case 1:
		return ""
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Forteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	}
	return "Zero"
}
