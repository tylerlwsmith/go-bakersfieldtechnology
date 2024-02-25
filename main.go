package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"bakersfieldtechnology.com/assets"
	"bakersfieldtechnology.com/components"
)

func main() {
	app := echo.New()
	assetHandler := assets.AssetHandler()

	app.GET("/assets/public/*", echo.WrapHandler(http.StripPrefix("/assets/public", assetHandler)))
	app.GET("/", func(c echo.Context) error {
		return components.Render(c, http.StatusOK, components.Homepage())
	})
	app.Logger.Fatal(app.Start(":3005"))
}
