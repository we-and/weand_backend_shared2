package utils

import "regexp"

func FilterStringKeepAlpha(s string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(s, "")
	return processedString, nil
}
func FilterStringKeepAlphaAndDot(s string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z.0-9]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(s, "")
	return processedString, nil
}
