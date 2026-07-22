package main

import "fmt"

func validateMonth(month int) bool {
	if (month < 1) || (month > 12) {
		fmt.Println("Month must be between 1 and 12")
		return false
	}
	return true
}

func validateYear(year int) bool {
	if year < 1975 {
		fmt.Println("Year must be after 1975!")
		return false
	}
	return true
}

func isLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

func daysInMonth(month, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 30
	}
}

func nameOfMonth(month int) string {
	switch month {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return "Invalid month"
	}
}

func main() {
	for {
		var month, year int
		fmt.Println("Enter month and year: ")
		fmt.Scan(&month, &year)
		if validateMonth(month) && validateYear(year) {
			monthName := nameOfMonth(month)
			fmt.Println("The number of days in", monthName, "of", year, "is", daysInMonth(month, year))
			break
		}
		fmt.Println("Invalid input. Try again:")
	}
}
