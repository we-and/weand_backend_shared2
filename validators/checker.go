package validators

func ValidateStrings(fields []ValidationFormat) (bool, string, string) {
	valid := true
	for _, field := range fields {
		isFieldValid, errorDesc := true, ""
		switch field.Type {
		case ValidationFormatKey_Email:
			isFieldValid, errorDesc = ValidateEmailField(field.Value)
			break
		case ValidationFormatKey_Nonempty:
			isFieldValid, errorDesc = ValidateNonEmpty(field.Value)
			break
		case ValidationFormatKey_MaxLength32:
			isFieldValid, errorDesc = ValidateMaxLengthString(field.Value, 32)
			break
		case ValidationFormatKey_Password:
			isFieldValid, errorDesc = ValidatePasswordString(field.Value)
			break
		case ValidationFormatKey_Alphabetical:
			isFieldValid, errorDesc = ValidateAlphabeticalString(field.Value)
			break
		case ValidationFormatKey_Number_NonEmpty:
			isFieldValid, errorDesc = ValidateNumberNonEmptyString(field.Value)
			break
		case ValidationFormatKey_Alphanumerical:
			isFieldValid, errorDesc = ValidateAlphanumericString(field.Value)
			break
		case ValidationFormatKey_Alphabetical_Nonempty:
			isFieldValid, errorDesc = ValidateAlphabeticalNonEmptyString(field.Value)
			break
		case ValidationFormatKey_Alphanumerical_Nonempty:
			isFieldValid, errorDesc = ValidateAlphanumericNonEmptyString(field.Value)
			break
		case ValidationFormatKey_Length3:
			isFieldValid, errorDesc = ValidateLengthString(field.Value, 3)
			break
		case ValidationFormatKey_MaxLength128:
			isFieldValid, errorDesc = ValidateMaxLengthString(field.Value, 128)
			break
		}
		valid = valid && isFieldValid
		if !valid {
			return false, field.KeyForDesc, errorDesc
		}
		//break
	}
	return valid, "", ""
}

func ValidateInts(fields []ValidationFormatInt) (bool, string, string) {
	valid := true
	for _, field := range fields {
		isFieldValid, errorDesc := true, ""
		switch field.Type {
		case ValidationFormatKey_NonZero:
			isFieldValid, errorDesc = ValidateNonZeroInt(field.Value)
			break
		}
		valid = valid && isFieldValid
		if !valid {
			return false, field.KeyForDesc, errorDesc
		}
		//break
	}
	return valid, "", ""
}
