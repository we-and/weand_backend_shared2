package business

import (
	"fmt"
)

func GetAmount(cur string, amount int32) string {
	if cur == "CAD" {
		damount:=float64(amount) /  100
		return fmt.Sprintf("$%.2f", damount)
	}else{
		damount:=float64(amount) /  100
		return fmt.Sprintf("%v%.2f",cur, damount)
	}
}
