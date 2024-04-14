package components

import (
	"bytes"
	"sync"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// https://stackoverflow.com/questions/30821745/idiomatic-way-to-handle-template-errors-in-golang
// https://blog.questionable.services/article/approximating-html-template-inheritance/
var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	b := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		b.Reset()
		bufferPool.Put(b)
	}()

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	err := t.Render(ctx.Request().Context(), ctx.Response().Writer)

	// I feel like there should be an `else` block that sets a 500 if there is
	// an error, but the app logs `superfluous response.WriteHeader call` when
	// when I do that.
	if err != nil {
		ctx.Response().Writer.WriteHeader(statusCode)
		b.WriteTo(ctx.Response().Writer)
	}

	return err
}
