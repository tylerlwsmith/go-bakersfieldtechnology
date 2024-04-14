package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"bakersfieldtechnology.com/assets"
	"bakersfieldtechnology.com/components"
	"bakersfieldtechnology.com/components/errorpage"
	"bakersfieldtechnology.com/components/homepage"
	"bakersfieldtechnology.com/components/privacypolicy"
)

func main() {
	app := echo.New()

	// https://echo.labstack.com/docs/error-handling#error-pages
	customHTTPErrorHandler := func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		if code == 404 {
			components.Render(c, 404, errorpage.NotFound())
			return
		}

		components.Render(c, code, errorpage.ServerError())
	}
	app.HTTPErrorHandler = customHTTPErrorHandler

	app.GET("/assets/public/*", echo.WrapHandler(http.StripPrefix("/assets/public", assets.AssetsHandler)))
	app.GET("/", func(c echo.Context) error {
		return components.Render(c, http.StatusOK, homepage.Homepage())
	})

	app.GET("/privacy-policy/", func(c echo.Context) error {
		return components.Render(c, http.StatusOK, privacypolicy.PrivacyPolicy())
	})

	// I tried to use echo's middleware.AddTrailingSlashWithConfig, but it
	// added a trailing slash to assets and broke URLs. Since this is the only
	// route that needs the slash, we'll handle it manually.
	app.GET("/privacy-policy", func(c echo.Context) error {
		return c.Redirect(301, "/privacy-policy/")
	})

	// https://pkg.go.dev/github.com/labstack/echo/v4@v4.11.4#Echo.StaticFS
	app.StaticFS("/*", assets.PublicFiles)

	app.RouteNotFound("/*", func(c echo.Context) error {
		return components.Render(c, http.StatusOK, homepage.Homepage())
	})

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = "3000"
	} else {
		port = strings.Trim(port, " ")
	}
	app.Logger.Fatal(app.Start(":" + port))
}
