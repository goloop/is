package is

import (
	"strings"
	"unicode"
)

// The numbers is a map of all Unicode code points that
// are classified as "Numbers".
var numbers = map[rune]struct{}{
	// Arabic numerals.
	'0': {}, '1': {}, '2': {}, '3': {}, '4': {},
	'5': {}, '6': {}, '7': {}, '8': {}, '9': {},

	// Fullwidth Arabic numerals.
	'０': {}, '１': {}, '２': {}, '３': {}, '４': {},
	'５': {}, '６': {}, '７': {}, '８': {}, '９': {},

	// Arabic-Indic numerals.
	'٠': {}, '١': {}, '٢': {}, '٣': {}, '٤': {},
	'٥': {}, '٦': {}, '٧': {}, '٨': {}, '٩': {},

	// Extended Arabic-Indic numerals.
	// Persian numerals.
	'۰': {}, '۱': {}, '۲': {}, '۳': {}, '۴': {},
	'۵': {}, '۶': {}, '۷': {}, '۸': {}, '۹': {},

	// Devanagari numerals.
	'०': {}, '१': {}, '२': {}, '३': {}, '४': {},
	'५': {}, '६': {}, '७': {}, '८': {}, '९': {},

	// Bengali numerals.
	'০': {}, '১': {}, '২': {}, '৩': {}, '৪': {},
	'৫': {}, '৬': {}, '৭': {}, '৮': {}, '৯': {},

	// Chinese numerals.
	'零': {}, '一': {}, '二': {}, '三': {}, '四': {},
	'五': {}, '六': {}, '七': {}, '八': {}, '九': {}, '十': {},
	'廿': {}, '卅': {}, '卌': {}, '百': {}, '千': {},
	'万': {}, '億': {}, '兆': {}, '京': {},
	'雲': {}, '穣': {},
	'拾': {}, '佰': {}, '仟': {}, '萬': {},
	'垓': {}, '秭': {}, '穰': {},

	// Japanese numerals.
	'〇': {}, '壱': {}, '弐': {}, '参': {},

	// Korean numerals.
	'영': {}, '일': {}, '이': {}, '삼': {}, '사': {},
	'오': {}, '육': {}, '칠': {}, '팔': {}, '구': {},
	'십': {}, '백': {}, '천': {}, '만': {},
	'억': {}, '조': {},

	// Gujarati numerals.
	'૦': {}, '૧': {}, '૨': {}, '૩': {}, '૪': {},
	'૫': {}, '૬': {}, '૭': {}, '૮': {}, '૯': {},

	// Punjabi (Gurmukhi) numerals.
	'੦': {}, '੧': {}, '੨': {}, '੩': {}, '੪': {},
	'੫': {}, '੬': {}, '੭': {}, '੮': {}, '੯': {},

	// Tamil numerals.
	'௦': {}, '௧': {}, '௨': {}, '௩': {}, '௪': {},
	'௫': {}, '௬': {}, '௭': {}, '௮': {}, '௯': {},

	// Telugu numerals.
	'౦': {}, '౧': {}, '౨': {}, '౩': {}, '౪': {},
	'౫': {}, '౬': {}, '౭': {}, '౮': {}, '౯': {},

	// Kannada numerals.
	'೦': {}, '೧': {}, '೨': {}, '೩': {}, '೪': {},
	'೫': {}, '೬': {}, '೭': {}, '೮': {}, '೯': {},

	// Malayalam numerals.
	'൦': {}, '൧': {}, '൨': {}, '൩': {}, '൪': {},
	'൫': {}, '൬': {}, '൭': {}, '൮': {}, '൯': {},

	// Thai numerals.
	'๐': {}, '๑': {}, '๒': {}, '๓': {}, '๔': {},
	'๕': {}, '๖': {}, '๗': {}, '๘': {}, '๙': {},

	// Lao numerals.
	'໐': {}, '໑': {}, '໒': {}, '໓': {}, '໔': {},
	'໕': {}, '໖': {}, '໗': {}, '໘': {}, '໙': {},

	// Tibetan numerals.
	'༠': {}, '༡': {}, '༢': {}, '༣': {}, '༤': {},
	'༥': {}, '༦': {}, '༧': {}, '༨': {}, '༩': {},

	// Myanmar (Burmese) numerals.
	'၀': {}, '၁': {}, '၂': {}, '၃': {}, '၄': {},
	'၅': {}, '၆': {}, '၇': {}, '၈': {}, '၉': {},

	// Khmer numerals.
	'០': {}, '១': {}, '២': {}, '៣': {}, '៤': {},
	'៥': {}, '៦': {}, '៧': {}, '៨': {}, '៩': {},

	// Hebrew numerals.
	'א': {}, 'ב': {}, 'ג': {}, 'ד': {},
	'ה': {}, 'ו': {}, 'ז': {}, 'ח': {}, 'ט': {},
	'י': {}, 'כ': {}, 'ל': {}, 'מ': {},
	'נ': {}, 'ס': {}, 'ע': {}, 'פ': {}, 'צ': {},
	'ק': {}, 'ר': {}, 'ש': {}, 'ת': {},
	'ך': {}, 'ם': {}, 'ן': {}, 'ף': {}, 'ץ': {},

	// Balinese numerals.
	'᭐': {}, '᭑': {}, '᭒': {}, '᭓': {}, '᭔': {},
	'᭕': {}, '᭖': {}, '᭗': {}, '᭘': {}, '᭙': {},

	// Limbu numerals.
	'᥆': {}, '᥇': {}, '᥈': {}, '᥉': {}, '᥊': {},
	'᥋': {}, '᥌': {}, '᥍': {}, '᥎': {}, '᥏': {},

	// Osmanya numerals.
	'𐒠': {}, '𐒡': {}, '𐒢': {}, '𐒣': {}, '𐒤': {},
	'𐒥': {}, '𐒦': {}, '𐒧': {}, '𐒨': {}, '𐒩': {},

	// Saurashtra numerals.
	'꣠': {}, '꣡': {}, '꣢': {}, '꣣': {}, '꣤': {},
	'꣥': {}, '꣦': {}, '꣧': {}, '꣨': {}, '꣩': {},

	// Sundanese numerals.
	'᮰': {}, '᮱': {}, '᮲': {}, '᮳': {}, '᮴': {},
	'᮵': {}, '᮶': {}, '᮷': {}, '᮸': {}, '᮹': {},

	// Javanese numerals.
	'꧐': {}, '꧑': {}, '꧒': {}, '꧓': {}, '꧔': {},
	'꧕': {}, '꧖': {}, '꧗': {}, '꧘': {}, '꧙': {},

	// Old Persian numerals.
	'𐏐': {}, '𐏑': {}, '𐏒': {}, '𐏓': {}, '𐏔': {},

	// Armenian numerals.
	'Ա': {}, 'Բ': {}, 'Գ': {}, 'Դ': {}, 'Ե': {}, 'Զ': {},
	'Է': {}, 'Ը': {}, 'Թ': {}, 'Ժ': {}, 'Ի': {}, 'Լ': {},
	'Խ': {}, 'Ծ': {}, 'Կ': {}, 'Հ': {}, 'Ձ': {}, 'Ղ': {},
	'Ճ': {}, 'Մ': {}, 'Յ': {}, 'Ն': {}, 'Շ': {}, 'Ո': {},
	'Չ': {}, 'Պ': {}, 'Ջ': {}, 'Ռ': {}, 'Ս': {}, 'Վ': {},
	'Տ': {}, 'Ր': {}, 'Ց': {}, 'Ւ': {}, 'Փ': {}, 'Ք': {},

	// Mongolian numerals.
	'᠑': {}, '᠒': {}, '᠓': {}, '᠔': {}, '᠕': {}, '᠖': {},
	'᠗': {}, '᠘': {}, '᠙': {}, '᠐': {}, 'ᠠ': {}, 'ᠡ': {},
	'ᠢ': {}, 'ᠣ': {}, 'ᠤ': {}, 'ᠥ': {},

	// Brahmi numerals.
	'𑁒': {}, '𑁓': {}, '𑁔': {}, '𑁕': {}, '𑁖': {}, '𑁗': {}, '𑁘': {},
	'𑁙': {}, '𑁚': {}, '𑁛': {}, '𑁜': {}, '𑁝': {}, '𑁞': {}, '𑁟': {},
	'𑁠': {}, '𑁡': {}, '𑁢': {}, '𑁣': {}, '𑁤': {}, '𑁥': {},

	// Ancient Greek numerals.
	'Α': {}, 'Β': {}, 'Γ': {}, 'Δ': {}, 'Ε': {}, 'Ϛ': {}, 'Ζ': {},
	'Η': {}, 'Θ': {}, 'Ι': {}, 'Κ': {}, 'Λ': {}, 'Μ': {}, 'Ν': {},
	'Ξ': {}, 'Ο': {}, 'Π': {}, 'Ϟ': {}, 'Ϡ': {}, 'ϡ': {}, 'Ϣ': {},
	'ϣ': {}, 'Ϥ': {}, 'ϥ': {}, 'Ϧ': {}, 'ϧ': {}, 'Ϩ': {}, 'ϩ': {},
	'Ϫ': {}, 'ϫ': {}, 'Ϭ': {}, 'ϭ': {}, 'Ϯ': {}, 'ϯ': {}, 'ϰ': {},
	'ϱ': {}, 'ϲ': {}, 'ϳ': {}, 'ϴ': {}, 'ϵ': {}, '϶': {}, 'Ϸ': {},
	'ϸ': {}, 'Ϻ': {}, 'ϻ': {},

	// Coptic numerals.
	'Ⲁ': {}, 'Ⲃ': {}, 'Ⲅ': {}, 'Ⲇ': {}, 'Ⲉ': {}, 'Ⲋ': {}, 'Ⲍ': {},
	'Ⲏ': {}, 'Ⲑ': {}, 'Ⲓ': {}, 'Ⲕ': {}, 'Ⲗ': {}, 'Ⲙ': {}, 'Ⲛ': {},
	'Ⲝ': {}, 'Ⲟ': {}, 'Ⲡ': {}, 'Ⲣ': {}, '⳰': {}, 'Ⳳ': {}, '⳴': {},
	'⳶': {}, '⳼': {}, '⳾': {}, '⳿': {},

	// Ethiopic numerals.
	'፩': {}, '፪': {}, '፫': {}, '፬': {}, '፭': {}, '፮': {},
	'፯': {}, '፰': {}, '፱': {}, '፲': {}, '፳': {}, '፴': {},
	'፵': {}, '፶': {}, '፷': {}, '፸': {}, '፹': {}, '፺': {},
	'፻': {}, '፼': {},
}

// Digit checks whether a string consists only of numbers.
//
// This method returns true if all characters in the string are numbers and
// the string is not empty. This includes digits (0-9), numeric characters
// that have a specific meaning in non-positional number systems (such as
// base 2, 8, or 16 number systems), and Unicode digit characters.
//
// Example usage:
//
//	is.Digit("1234")     // Output: true
//	is.Digit("Ⅳ")       // Output: false
//	is.Digit("1234abc")  // Output: false
func Digit(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Numeric is similar to Digit, but also recognizes numeric characters
// used in other number systems.
//
// Returns true if all characters in the string are any numeric and
// false otherwise. It recognizes numbers (including numbers with a
// decimal point), Roman numerals, Chinese numerals, and more.
//
// Example usage:
//
//	is.Numeric("1234")     // Output: true
//	is.Numeric("Ⅳ")        // Output: true
//	is.Numeric("1234abc")  // Output: false
func Numeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Check for chinese numerals and regular digits.
	for _, r := range s {
		if !unicode.Is(unicode.Number, r) {
			if _, ok := numbers[r]; !ok {
				return false
			}
		}
	}
	return true
}

// Decimal returns true if all characters in the string are decimal digits
// and the string is not empty. Decimal digits are only the characters with
// the numbers 0 through 9. It does not recognize any other numeric characters
// such as Roman numerals or digits from non-positional number systems.
//
// Example usage:
//
//	is.Decimal("1234")     // Output: true
//	is.Decimal("Ⅳ")       // Output: false
//	is.Decimal("1234abc")  // Output: false
//	is.Decimal("1234.56")  // Output: false
func Decimal(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

// Alpha checks whether a string consists only of alphabetic
// characters (letters). It does not recognize digits, special
// characters.
//
// Example usage:
//
//	is.Alpha("Київ")  // Output: true
//	is.Alpha("abc")   // Output: true
//	is.Alpha("abc1")  // Output: false, contains a digit
//	is.Alpha("abc!")  // Output: false, contains a special character
//	is.Alpha("abcΔ")  // Output: false, contains a non-Latin letter
func Alpha(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Alnum checks whether a string consists only of alphabetic
// characters (letters) and numbers.
//
// Example usage:
//
//	is.Alnum("abc123")      // Output: true
//	is.Alnum("abc")         // Output: true
//	is.Alnum("123")         // Output: true
//	is.Alnum("abc!")        // Output: false
//	is.Alnum("abcΔ")        // Output: false
func Alnum(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

// Lower checks whether a string consists only of lowercase
// alphabetic characters.
//
// Example usage:
//
//	is.Lower("abc")       // Output: true
//	is.Lower("Abc")       // Output: false, contains an uppercase letter
//	is.Lower("abc123")    // Output: false, contains a number
//	is.Lower("abc!")      // Output: false, contains a special character
func Lower(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

// Upper checks whether a string consists only of uppercase
// alphabetic characters.
//
// Example usage:
//
//	is.Upper("ABC")      // Output: true
//	is.Upper("AbC")      // Output: false, contains a lowercase letter
//	is.Upper("ABC123")   // Output: false, contains a number
//	is.Upper("ABC!")     // Output: false, contains a special character
func Upper(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// Title checks whether a string is a titlecased string.
//
// In a titlecased string, upper- and title-case characters may only
// follow uncased characters and lowercase characters only cased ones.
//
// Example usage:
//
//	is.Title("Hello World")   // Output: true
//	is.Title("Hello world")   // Output: false, 'world' starts with a lowercase
//	is.Title("HELLO WORLD")   // Output: false, all letters are uppercase
//	is.Title("hELLO wORLD")   // Output: false, words start with a lowercase
func Title(s string) bool {
	if len(s) == 0 {
		return false
	}

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	for _, word := range words {
		for index, char := range word {
			if index == 0 {
				if !unicode.IsUpper(char) {
					return false
				}
			} else {
				if !unicode.IsLower(char) {
					return false
				}
			}
		}
	}
	return true
}

// Space checks whether a string consists only of whitespace characters.
// Whitespace characters includes spaces, tabs, newlines, and other
// Unicode whitespace characters.
//
// Example usage:
//
//	is.Space(" \t\n")    // Output: true
//	is.Space("Hello")    // Output: false
//	is.Space(" ")        // Output: true
//	is.Space("\n\t ")    // Output: true
//	is.Space("")         // Output: false, an empty string has no characters
func Space(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}
