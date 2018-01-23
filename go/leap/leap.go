// Package leap test if given year is a leap year or common year
package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	if year%4 == 0 { //on every year that is evenly divisible by 4
		if year%100 == 0 { //except every year that is evenly divisible by 100
			if year%400 == 0 { //unless the year is also evenly divisible by 400
				return true
			}
			return false
		}
		return true
	}
	return false
}
