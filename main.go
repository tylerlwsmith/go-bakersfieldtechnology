package main

import (
	"net/http"
	"os"
	"strings"

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

	app.GET("/privacy-policy/", func(c echo.Context) error {
		return components.Render(c, http.StatusOK, components.PrivacyPolicy())
	})

	// TODO: should I export a handler for public files? There would be more
	// symmetry with the AssetHandler if I did.
	app.GET("/*", echo.WrapHandler(http.FileServer(http.FS(assets.PublicFiles))))

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = "3000"
	} else {
		port = strings.Trim(port, " ")
	}
	app.Logger.Fatal(app.Start(":" + port))
}
