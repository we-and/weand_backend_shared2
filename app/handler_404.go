package app

import (
	"errors"
	response "stretches-common-api/response"

	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		response.SetNotFound(c, "404", errors.New("Not found"))
		return nil
	}
}
