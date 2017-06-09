package helpers

import "strconv"

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

func padTimes(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func PadLeft(str string, length int, pad string) string {
	return padTimes(pad, length-len(str)) + str
}

// Right right-pads the string with pad up to len runes
func PadRight(str string, length int, pad string) string {
	return str + padTimes(pad, length-len(str))
}

func PadIntLeft(num int, len int, pad string) string {
	str := strconv.Itoa(num)
	return PadLeft(str, len, pad)
}
