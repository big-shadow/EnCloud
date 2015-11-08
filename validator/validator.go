package validator

import (
	"EnCloud/logger"
	"regexp"
)

func ValidateEmailAddress(email string) bool {

	if len([]rune(email)) < 5 {
		return false
	}

	// Regex used ASP.NET's RegularExpressionValidator
	match, _ := regexp.MatchString("^\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", email)

	if r := recover(); r != nil {
		logger.Log("Recovered in validator.go : ValidateEmailAddress(email string)")
	}

	return match
}

func ValidatePassword(email string) bool {

	// Make some complexity rules.

	return true
}
