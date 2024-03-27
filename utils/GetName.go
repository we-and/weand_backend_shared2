package utils

import (
	"fmt"
)

func GetName(FirstName string, MiddleName string, LastName string) string {
	if len(MiddleName) > 0 {
		return fmt.Sprintf("%v %v %v", FirstName, MiddleName, LastName)
	} else {
		return fmt.Sprintf("%v %v", FirstName, LastName)
	}
}

func GetAbbreviatedName(FirstName string, MiddleName string, LastName string) string {
	if len(FirstName) > 0 && len(MiddleName) > 0 && len(LastName) > 0 {
		return fmt.Sprintf("%v %v. %v.", FirstName, MiddleName[0:1], LastName[0:1])
	} else if len(FirstName) > 0 && len(LastName) > 0 {
		return fmt.Sprintf("%v %v.", FirstName, LastName[0:1])
	} else if len(FirstName) > 0 {
		return fmt.Sprintf("%v", FirstName)
	} else if len(LastName) > 0 {
		return fmt.Sprintf("%v", LastName)
	} else {
		return fmt.Sprintf("Undefined")
	}
}
