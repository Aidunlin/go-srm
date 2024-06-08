package value

import (
	"fmt"
	"time"
)

func GetDegrees() []string {
	return []string{
		"Advanced Manufacturing",
		"Business Administration",
		"Cuisine Management",
		"Cybersecurity",
		"Digital Media Arts",
		"Fine Arts",
		"Management",
		"Marketing",
		"Music",
		"Network Technology",
		"Professional Baking and Pastry Arts",
		"Web Development",
	}
}

func IsDegreeProgram(value string) bool {
	if value == "" {
		return true
	}

	for _, degree := range GetDegrees() {
		if value == degree {
			return true
		}
	}

	return false
}

func IsGraduationDate(value string) bool {
	if value == "" {
		return true
	}

	_, err := time.Parse(time.DateOnly, value)
	return err == nil
}

func IsFinancialAid(value int) bool {
	return value == 0 || value == 1
}

func DisplayDate(value string) string {
	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%02d/%02d/%04d", date.Month(), date.Day(), date.Year())
}
