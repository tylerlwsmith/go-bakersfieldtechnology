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

	// TODO: should I export a handler for public files? There would be more
	// symmetry with the AssetHanlder if I did.
	app.GET("/*", echo.WrapHandler(http.FileServer(http.FS(assets.PublicFiles))))
	app.Logger.Fatal(app.Start(":3005"))
}
