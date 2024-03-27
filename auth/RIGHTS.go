package auth

import (
	"strings"

)

func HasAdminRights(rights string) bool{
		return checkHasRights(rights,"ADMIN") || checkHasRights(rights,"TEMPADMIN")

}
func HasUserRights(rights string) bool {
	return checkHasRights(rights,"USER")
}
func checkHasRights(rights string,needed string) bool {
	parts:=	strings.Split(rights,"-")
//	fmt.Printf("%v",parts )
	for  _, part :=range parts{
		if part == needed {
			return true
		}
	}
	return false
}
