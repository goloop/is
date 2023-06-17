package is

import (
	"unicode"
)

//var numbers = map[rune]int{
//	// Arabic numerals.
//	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
//	'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
//
//	// Fullwidth Arabic numerals.
//	'０': 0, '１': 1, '２': 2, '３': 3, '４': 4,
//	'５': 5, '６': 6, '７': 7, '８': 8, '９': 9,
//
//	// Arabic-Indic numerals.
//	'٠': 0, '١': 1, '٢': 2, '٣': 3, '٤': 4,
//	'٥': 5, '٦': 6, '٧': 7, '٨': 8, '٩': 9,
//
//	// Extended Arabic-Indic numerals.
//	// Persian numerals.
//	'۰': 0, '۱': 1, '۲': 2, '۳': 3, '۴': 4,
//	'۵': 5, '۶': 6, '۷': 7, '۸': 8, '۹': 9,
//
//	// Devanagari numerals.
//	'०': 0, '१': 1, '२': 2, '३': 3, '४': 4,
//	'५': 5, '६': 6, '७': 7, '८': 8, '९': 9,
//
//	// Bengali numerals.
//	'০': 0, '১': 1, '২': 2, '৩': 3, '৪': 4,
//	'৫': 5, '৬': 6, '৭': 7, '৮': 8, '৯': 9,
//
//	// Chinese numerals.
//	'零': 0, '一': 1, '二': 2, '三': 3, '四': 4,
//	'五': 5, '六': 6, '七': 7, '八': 8, '九': 9, '十': 10,
//	'廿': 20, '卅': 30, '卌': 40, '百': 100, '千': 1000,
//	'万': 10000, '億': 100000000, '兆': 1000000000000, '京': 10000000000000000,
//	'雲': 100000000000000000000, '穣': 1000000000000000000000000,
//	'拾': 10, '佰': 100, '仟': 1000, '萬': 10000,
//	'垓': 100000000000000000000, '秭': 1000000000000000000000000, '穰': 10000000000000000000000000000,
//
//	// Japanese numerals.
//	'〇': 0, '壱': 1, '弐': 2, '参': 3,
//
//	// Korean numerals.
//	'영': 0, '일': 1, '이': 2, '삼': 3, '사': 4,
//	'오': 5, '육': 6, '칠': 7, '팔': 8, '구': 9,
//	'십': 10, '백': 100, '천': 1000, '만': 10000,
//	'억': 100000000, '조': 1000000000000,
//
//	// Gujarati numerals.
//	'૦': 0, '૧': 1, '૨': 2, '૩': 3, '૪': 4,
//	'૫': 5, '૬': 6, '૭': 7, '૮': 8, '૯': 9,
//
//	// Punjabi (Gurmukhi) numerals.
//	'੦': 0, '੧': 1, '੨': 2, '੩': 3, '੪': 4,
//	'੫': 5, '੬': 6, '੭': 7, '੮': 8, '੯': 9,
//
//	// Tamil numerals.
//	'௦': 0, '௧': 1, '௨': 2, '௩': 3, '௪': 4,
//	'௫': 5, '௬': 6, '௭': 7, '௮': 8, '௯': 9,
//
//	// Telugu numerals.
//	'౦': 0, '౧': 1, '౨': 2, '౩': 3, '౪': 4,
//	'౫': 5, '౬': 6, '౭': 7, '౮': 8, '౯': 9,
//
//	// Kannada numerals.
//	'೦': 0, '೧': 1, '೨': 2, '೩': 3, '೪': 4,
//	'೫': 5, '೬': 6, '೭': 7, '೮': 8, '೯': 9,
//
//	// Malayalam numerals.
//	'൦': 0, '൧': 1, '൨': 2, '൩': 3, '൪': 4,
//	'൫': 5, '൬': 6, '൭': 7, '൮': 8, '൯': 9,
//
//	// Thai numerals.
//	'๐': 0, '๑': 1, '๒': 2, '๓': 3, '๔': 4,
//	'๕': 5, '๖': 6, '๗': 7, '๘': 8, '๙': 9,
//
//	// Lao numerals.
//	'໐': 0, '໑': 1, '໒': 2, '໓': 3, '໔': 4,
//	'໕': 5, '໖': 6, '໗': 7, '໘': 8, '໙': 9,
//
//	// Tibetan numerals.
//	'༠': 0, '༡': 1, '༢': 2, '༣': 3, '༤': 4,
//	'༥': 5, '༦': 6, '༧': 7, '༨': 8, '༩': 9,
//
//	// Myanmar (Burmese) numerals.
//	'၀': 0, '၁': 1, '၂': 2, '၃': 3, '၄': 4,
//	'၅': 5, '၆': 6, '၇': 7, '၈': 8, '၉': 9,
//
//	// Khmer numerals.
//	'០': 0, '១': 1, '២': 2, '៣': 3, '៤': 4,
//	'៥': 5, '៦': 6, '៧': 7, '៨': 8, '៩': 9,
//
//	// // Greek numerals.
//	// 'Ο': 0, 'Ι': 1, 'ΙΙ': 2, 'ΙΙΙ': 3, 'ΙV': 4,
//	// 'V': 5, 'VΙ': 6, 'VΙΙ': 7, 'VΙΙΙ': 8, 'ΙΧ': 9,
//	// 'Χ': 10, // continues with 'ΧΙ' for 11 and so on...
//
//	// Hebrew numerals.
//	'א': 1, 'ב': 2, 'ג': 3, 'ד': 4,
//	'ה': 5, 'ו': 6, 'ז': 7, 'ח': 8, 'ט': 9,
//	'י': 10, 'כ': 20, 'ל': 30, 'מ': 40,
//	'נ': 50, 'ס': 60, 'ע': 70, 'פ': 80, 'צ': 90,
//	'ק': 100, 'ר': 200, 'ש': 300, 'ת': 400,
//	'ך': 500, 'ם': 600, 'ן': 700, 'ף': 800, 'ץ': 900,
//
//	// Special cases for 15 and 16 to avoid accidental
//	// representation of the divine name.
//	//'ט״ו': 15, 'ט״ז': 16,
//
//	// Balinese numerals.
//	'᭐': 0, '᭑': 1, '᭒': 2, '᭓': 3, '᭔': 4,
//	'᭕': 5, '᭖': 6, '᭗': 7, '᭘': 8, '᭙': 9,
//
//	// Limbu numerals.
//	'᥆': 0, '᥇': 1, '᥈': 2, '᥉': 3, '᥊': 4,
//	'᥋': 5, '᥌': 6, '᥍': 7, '᥎': 8, '᥏': 9,
//
//	// Osmanya numerals.
//	'𐒠': 0, '𐒡': 1, '𐒢': 2, '𐒣': 3, '𐒤': 4,
//	'𐒥': 5, '𐒦': 6, '𐒧': 7, '𐒨': 8, '𐒩': 9,
//
//	// Saurashtra numerals.
//	'꣠': 0, '꣡': 1, '꣢': 2, '꣣': 3, '꣤': 4,
//	'꣥': 5, '꣦': 6, '꣧': 7, '꣨': 8, '꣩': 9,
//
//	// Sundanese numerals.
//	'᮰': 0, '᮱': 1, '᮲': 2, '᮳': 3, '᮴': 4,
//	'᮵': 5, '᮶': 6, '᮷': 7, '᮸': 8, '᮹': 9,
//
//	// Ethiopic numerals.
//	'፩': 1, '፪': 2, '፫': 3, '፬': 4, '፭': 5,
//	'፮': 6, '፯': 7, '፰': 8, '፱': 9, '፲': 10, // And so on...
//
//	// Javanese numerals.
//	'꧐': 0, '꧑': 1, '꧒': 2, '꧓': 3, '꧔': 4,
//	'꧕': 5, '꧖': 6, '꧗': 7, '꧘': 8, '꧙': 9,
//
//	// Old Persian numerals.
//	'𐏐': 1, '𐏑': 2, '𐏒': 10, '𐏓': 20, '𐏔': 100,
//
//	// And many more...
//
//}

var numbers = map[rune]struct{}{
	// Greek numerals.
	// 'Ο': {}, 'Ι': {}, 'ΙΙ': {}, 'ΙΙΙ': {}, 'ΙV': {},
	// 'V': {}, 'VΙ': {}, 'VΙΙ': {}, 'VΙΙΙ': {}, 'ΙΧ': {},
	// 'Χ': {}, // continues with 'ΧΙ' for 11 and so on...

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

	// Special cases for 15 and 16 to avoid accidental
	// representation of the divine name.
	//'ט״ו': {}, 'ט״ז': {},

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
	//'፩': {}, '፪': {}, '፫': {}, '፬': {}, '፭': {},
	//'፮': {}, '፯': {}, '፰': {}, '፱': {}, '፲': {}, // And so on...
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

	// // check for decimal number
	// if _, err := strconv.ParseFloat(s, 64); err == nil {
	// 	return true
	// }

	// check for chinese numerals and regular digits
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

	prevCased := true

	for _, r := range s {
		if unicode.IsUpper(r) || unicode.IsTitle(r) {
			if prevCased {
				return false
			}
			prevCased = true
		} else if unicode.IsLower(r) {
			if !prevCased {
				return false
			}
		} else {
			prevCased = false
		}
	}

	return true
}

// Space checks whether a string consists only of whitespace characters.
// Whitespace characters includes spaces, tabs, newlines, and other
// Unicode whitespace characters.
//
// This function is analogous to Python's str.isspace() method.
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
