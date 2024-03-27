package utils

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func RemoveFromStringArray(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
