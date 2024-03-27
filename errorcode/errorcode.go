package errorcode

import "fmt"

func Build(routeCode string, queryCode string) string {
	return fmt.Sprintf("%v-%v", routeCode, queryCode)
}
