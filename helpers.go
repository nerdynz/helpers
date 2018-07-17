package helpers

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

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

func Round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

// ToFixed rounds to the specified decimal place
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

func KebabCase(in string) string {
	out := SnakeCase(in)
	out = strings.Replace(out, "_", "-", -1)
	return out
}

func SnakeCase(in string) string {
	in = strings.Replace(in, " ", "_", -1)
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return strings.Replace(string(out), "__", "_", -1)
}

func SplitTitleCase(in string) string {
	out := titleCase(in, " ")
	return out
}

func TitleCase(in string) string {
	out := titleCase(in, "")
	return out
}

func titleCase(in string, sep string) string {
	out := SnakeCase(in)
	words := strings.Split(out, "_")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	out = strings.Join(words, sep)
	return out
}

func Substring(str string, length int) string {
	if len(str) > length {
		return strings.Join(strings.Split(str, "")[:length], "")
	}
	return str
}
