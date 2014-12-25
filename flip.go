package flip

import "strings"

var upsideDownChars = map[string]string{
	"z": "z", "y": "ʎ", "x": "x", "w": "ʍ", "v": "ʌ", "u": "n", "t": "ʇ",
	"s": "s", "r": "ɹ", "q": "b", "p": "d", "o": "o", "n": "u", "m": "ɯ",
	"l": "ן", "k": "ʞ", "j": "ɾ", "i": "ı", "h": "ɥ", "g": "ƃ", "f": "ɟ",
	"e": "ǝ", "d": "p", "c": "ɔ", "b": "q", "a": "ɐ", " ": " ", "-": "-",
	"+": "+", "Z": "Z", "Y": "⅄", "X": "X", "W": "M", "V": "Λ", "U": "∩",
	"T": "┴", "S": "S", "R": "ɹ", "Q": "Q", "P": "Ԁ", "O": "O", "N": "N",
	"M": "W", "L": "˥", "K": "ʞ", "J": "ſ", "I": "I", "H": "H", "G": "פ",
	"F": "Ⅎ", "E": "Ǝ", "D": "p", "C": "Ɔ", "B": "q", "A": "∀", "9": "6",
	"8": "8", "7": "ㄥ", "6": "9", "5": "ϛ", "4": "ㄣ", "3": "Ɛ",
	"2": "ᄅ", "1": "Ɩ", "0": "0",
}

// UpsideDown turns the given string upside down
func UpsideDown(s string) string {
	ss := strings.Split(s, "")
	ns := ""

	for i := len(ss) - 1; i >= 0; i-- {
		uc := upsideDownChars[ss[i]]

		if uc == "" {
			uc = upsideDownChars[strings.ToLower(ss[i])]
		}

		if uc == "" {
			uc = ""
		}

		ns += uc
	}

	return ns
}

// Table flips the table
func Table(s string) string {
	return "(╯°□°）╯︵" + UpsideDown(s)
}

// Reverse reverses the given string
func Reverse(s string) string {
	ss := strings.Split(s, "")
	ns := ""

	for i := len(ss) - 1; i >= 0; i-- {
		ns += ss[i]
	}

	return ns
}
