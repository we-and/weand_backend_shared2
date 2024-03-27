package validators

import (
	"fmt"
	"regexp"
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"
const alphanumeric = "abcdefghijklmnopqrstuvwxyz0123456789"
const numeric = "0123456789"
const number = "0123456789.,+-"

func ValidateAlphanumericString(field string) (bool, string) {
	if !isAlphanumericOnly(field) {
		return false, "Invalid parameter not alphanumerical"
	}
	return true, ""
}
func ValidateAlphanumericNonEmptyString(field string) (bool, string) {
	if !isAlphanumericOnly(field) {
		return false, "Invalid parameter not alphanumerical"
	}
	if !isNonEmpty(field) {
		return false, ("Invalid parameter not non empty")
	}
	return true, ""
}
func ValidateNonEmptyString(field string) (bool, string) {
	if !isNonEmpty(field) {
		return false, ("Invalid parameter not non empty")
	}
	return true, ""
}

func ValidateNumericalString(field string) (bool, string) {
	if !isNumericOnly(field) {
		return false, ("Invalid parameter not alphabetical")
	}
	return true, ""
}
func ValidateIdString(field string) (bool, string) {
	if !isNumericOnly(field) {
		return false, ("Invalid parameter not an ID")
	}
	return true, ""
}
func ValidatePidString(field string) (bool, string) {
	if !isNumericOnly(field) {
		return false, ("Invalid parameter not a pid")
	}
	return true, ""
}

func isAlphabeticalOnly(s string) bool {
	return checkStringCharacters(s, alpha)
}
func isNonMaxLength(s string, maxLength int) bool {
	return len(s) < maxLength
}
func isNumberOnly(s string) bool {
	return checkStringCharacters(s, number)
}
func isNonEmpty(s string) bool {
	return len(s) > 0
}
func isNumericOnly(s string) bool {
	return checkStringCharacters(s, numeric)
}
func isAlphanumericOnly(s string) bool {
	return checkStringCharacters(s, alphanumeric)
}
func checkStringCharacters(s string, whiteList string) bool {
	for _, char := range s {
		if !strings.Contains(whiteList, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
func ValidateAlphabeticalString(field string) (bool, string) {
	if !isAlphabeticalOnly(field) {
		return false, ("Invalid parameter not alphabetical")
	}
	return true, ""
}
func ValidateMaxLengthString(field string, maxLength int) (bool, string) {
	if !isNonMaxLength(field, maxLength) {
		return false, fmt.Sprintf("Invalid parameter: longer than %v characters", maxLength)
	}
	return true, ""
}
func ValidateLengthString(field string, length int) (bool, string) {
	if len(field) != length {
		return false, fmt.Sprintf("Invalid parameter: wrong length should be %v characters", length)
	}
	return true, ""
}
func ValidatePasswordString(field string) (bool, string) {
	if !isNonMaxLength(field, 32) {
		return false, fmt.Sprintf("Invalid parameter: longer than %v characters", 32)
	}
	return true, ""
}
func ValidateAlphabeticalNonEmptyString(field string) (bool, string) {
	if !isAlphabeticalOnly(field) {
		return false, ("Invalid parameter: not alphabetical ")
	}
	if !isNonEmpty(field) {
		return false, ("Invalid parameter: not non empty")
	}
	return true, ""
}
func ValidateNumberNonEmptyString(field string) (bool, string) {
	if !isNumberOnly(field) {
		return false, ("Invalid parameter: not number ")
	}
	if !isNonEmpty(field) {
		return false, ("Invalid parameter: not non empty")
	}
	return true, ""
}

func ValidateEmailField(email string) (bool, string) {
	field := email
	if len(field) == 0 {
		return false, ("Invalid parameter")
	}

	var rxEmail = regexp.MustCompile(".+@.+\\..+")
	if !rxEmail.Match([]byte(field)) {
		return false, "Invalid parameter"
	}
	return true, ""
}

func ValidateNonEmpty(field string) (bool, string) {

	if len(field) == 0 {
		return false, ("Invalid parameter: not nonempty")
	}
	return true, ""
}

func ValidateNonZeroInt(field uint32) (bool, string) {
	if field == 0 {
		return false, ("Invalid parameter: not nonempty")
	}
	return true, ""
}

func validateField(fieldvalue string) (bool, string) {
	if len(fieldvalue) == 0 {
		return false, "Invalid parameter"
	}

	return true, ""

}
