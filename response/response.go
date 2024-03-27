package response

import (
	"fmt"
	"stretches-common-api/types"

	"encoding/xml"

	"github.com/gofiber/fiber/v2"
)

func SetUnauthorized(c *fiber.Ctx, trigger string, err error, errCode string) error {
	route := c.OriginalURL()
	//log.MLogger.Error.Println("UNAUTH")
	c.Status(fiber.StatusForbidden).JSON(types.ErrorMsg{
		Level:   "FORBIDDEN",
		Route:   route,
		Trigger: trigger,
		ErrCode: errCode,
		Desc:    fmt.Sprintf("%v", err)})
	return nil
}

func SetNotFound(c *fiber.Ctx, trigger string, err error) {
	route := c.OriginalURL()

	//RETURN ERROR
	c.Status(fiber.StatusInternalServerError).JSON(types.ErrorMsg{
		Level:   "WARN",
		Route:   route,
		Trigger: trigger,
		Desc:    fmt.Sprintf("%v", err)})

}
func SetBadRequest(c *fiber.Ctx, trigger string, err error, errCode string) {

	route := c.OriginalURL()
	c.Status(fiber.StatusBadRequest).JSON(types.ErrorMsg{
		Level:   "WARN",
		Route:   route,
		Trigger: trigger,
		ErrCode: errCode,
		Desc:    fmt.Sprintf("%v", err)})
}

func SetInternalError(c *fiber.Ctx, trigger string, err error, errCode string) {
	route := c.OriginalURL()
	//RETURN ERROR
	c.Status(fiber.StatusInternalServerError).JSON(types.ErrorMsg{
		Level:   "ERROR",
		Route:   route,
		Trigger: trigger,
		Desc:    fmt.Sprintf("%v", err),
		ErrCode: errCode,
	})
}

func SetOK(c *fiber.Ctx, obj interface{}) {
	c.Status(fiber.StatusOK).JSON(obj)
}

func SetOKXML(c *fiber.Ctx, obj interface{}) {
	c.Set("Content-type", "text/xml")
	c.Status(fiber.StatusOK)
	raw, err := xml.Marshal(obj)
	if err != nil {
		fmt.Errorf("error serializing xml: %v", obj)
	}
	header := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	withHeader := fmt.Sprintf("%v\n%v", header, string(raw))
	c.Send([]byte(withHeader))
	// c.XML(obj)
}

func SetOKHTML(c *fiber.Ctx, obj string) {
	c.Set("Content-type", "text/html")
	c.Status(fiber.StatusOK)
	header := "<html encoding=\"UTF-8\"?>"
	footer := "</html>"
	withHeader := fmt.Sprintf("%v\n%v\n%v", header, obj, footer)
	c.Send([]byte(withHeader))
	// c.XML(obj)
}
