package utils

import "time"

// CalculateAge returns the age in years for the provided date of birth.
// It uses the current time and handles birthdays that have not occurred yet
// in the current year, including leap year dates. It never returns a negative age.
func CalculateAge(dob time.Time) int {
	if dob.IsZero() {
		return 0
	}

	now := time.Now()
	age := now.Year() - dob.Year()

	// If birthday hasn't occurred yet this year, subtract one year.
	birthdayThisYear := time.Date(now.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, now.Location())
	if now.Before(birthdayThisYear) {
		age--
	}

	if age < 0 {
		return 0
	}
	return age
}
 