// Package leap provides a function for determing whether or not a given year is a leap year
package leap

// IsLeapYear takes an int year and returns whether or not that year is a leap year
func IsLeapYear(num int) bool {
	if num%400 == 0 {
		return true
	}

	if num%4 == 0 && num%100 == 0 {
		return false
	}

	if num%4 == 0 {
		return true
	}
	return false
}
