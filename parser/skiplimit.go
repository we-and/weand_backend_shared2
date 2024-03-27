package parser
import (
	"errors"
	fiber "github.com/gofiber/fiber/v2"

	"strconv"
)
func ParseSkipLimit(c *fiber.Ctx) (int,int,string,error,string ){
	//PARSE skip limit
	skipStr := c.Params("skip")
	limitStr := c.Params("limit")

	skip, errSkip := strconv.Atoi(skipStr)
	if errSkip != nil {
		return -1,-1, "Validation failed for field skip", errors.New(""), "ME01420"
	}
	limit, errLimit := strconv.Atoi(limitStr)
	if errLimit != nil {
		return -1,-1,"Validation failed for field limit", errors.New(""), "ME01421"
	}
	return skip,limit,"",nil,""
}