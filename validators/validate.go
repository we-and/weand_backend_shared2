package validators

import (
	"fmt"
)

type ValidationFormatKey string

func GetEmailValidator(email string) []ValidationFormat {
	return []ValidationFormat{GetValidatorFormat(email, "email", ValidationFormatKey_Email)}
}
func GetCurrencyValidator(field string) []ValidationFormat {
	return []ValidationFormat{GetValidatorFormat(field, "currency", ValidationFormatKey_Length3)}
}
func GetNameValidator(field string) []ValidationFormat {
	return []ValidationFormat{GetValidatorFormat(field, "name", ValidationFormatKey_MaxLength128)}
}
func GetIdValidator(field string) []ValidationFormat {
	return []ValidationFormat{GetValidatorFormat(field, "name", ValidationFormatKey_Nonempty)}
}

func GetStringsNonEmptyValidator(fields []string) []ValidationFormat {
	res := []ValidationFormat{}
	for k, v := range fields {
		res = append(res, GetValidatorFormat(v, fmt.Sprintf("Field %v", k), ValidationFormatKey_Nonempty))
	}
	return res
}

func GetEmailConfirmTokenValidator(token string) []ValidationFormat {
	return []ValidationFormat{
		GetValidatorFormat(token, "token", ValidationFormatKey_Alphanumerical),
		GetValidatorFormat(token, "token", ValidationFormatKey_MinLength8)}
}
func GetPasswordValidator(password string) []ValidationFormat {
	return []ValidationFormat{GetValidatorFormat(password, "password", ValidationFormatKey_MaxLength32)}
}
func GetValidatorFormat(value string, key string, type_ ValidationFormatKey) ValidationFormat {
	return ValidationFormat{
		Type:       type_,
		Value:      value,
		KeyForDesc: key,
	}
}

func GetValidatorFormatInt(value uint32, key string, type_ ValidationFormatKey) ValidationFormatInt {
	return ValidationFormatInt{
		Type:       type_,
		Value:      value,
		KeyForDesc: key,
	}
}
