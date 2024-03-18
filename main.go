package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"bakersfieldtechnology.com/assets"
	"bakersfieldtechnology.com/components"
	"bakersfieldtechnology.com/components/homepage"
)

func main() {
	app := echo.New()

	app.GET("/assets/public/*", echo.WrapHandler(http.StripPrefix("/assets/public", assets.AssetsHandler)))
	app.GET("/", func(c echo.Context) error {
		return components.Render(c, http.StatusOK, homepage.Homepage())
	})
	app.Logger.Fatal(app.Start(":3005"))
}
