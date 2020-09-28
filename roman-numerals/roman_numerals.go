// Package romannumerals provides method for converting an integer into a roman numeral
package romannumerals

import "fmt"

// ToRomanNumeral takes an integer and returns the value converted to a roman numeral string
func ToRomanNumeral(num int) (string, error) {
	romanNumeral := ""

	if num <= 0 || num >= 3001 {
		return romanNumeral, fmt.Errorf("numeral must be greater than 0 and less than 3001. value given %d", num)
	}

	for num > 0 {
		switch {
		case num >= 1000:
			romanNumeral += "M"
			num -= 1000
		case num >= 900:
			romanNumeral += "CM"
			num -= 900
		case num >= 500:
			romanNumeral += "D"
			num -= 500
		case num >= 400:
			romanNumeral += "CD"
			num -= 400
		case num >= 100:
			romanNumeral += "C"
			num -= 100
		case num >= 90:
			romanNumeral += "XC"
			num -= 90
		case num >= 50:
			romanNumeral += "L"
			num -= 50
		case num >= 40:
			romanNumeral += "XL"
			num -= 40
		case num >= 10:
			romanNumeral += "X"
			num -= 10
		case num >= 9:
			romanNumeral += "IX"
			num -= 9
		case num >= 5:
			romanNumeral += "V"
			num -= 5
		case num >= 4:
			romanNumeral += "IV"
			num -= 4
		case num >= 1:
			romanNumeral += "I"
			num -= 1
		}
	}

	return romanNumeral, nil
}
