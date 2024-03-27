package common

import (
	"fmt"
	"github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"

)

///create routes for api documentation
///this will give access to a swagger file when requesting /swagger

func ApplyDocumentation(app *fiber.App, path string ) {
	app.Static("/docs", "./docs")
	app.Use("/swagger", swagger.New(swagger.Config{ // custom
		URL:        fmt.Sprintf( "%v/docs/swagger.json",path),
		DeepLinking: true,
	}))

}
