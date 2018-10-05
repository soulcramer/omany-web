package pages

import (
	"github.com/aerogo/aero"
)

// Configure registers the page routes in the application.
func Configure(app *aero.Application) {
	app.Get("/", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}
