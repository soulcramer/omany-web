package main

import (
	"./db"
	"./pages"
	"./utils/routetests"
	"github.com/aerogo/aero"
	"github.com/aerogo/session-store-nano"
)

var app = aero.New()

func main() {
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {

	app.Sessions.Duration = 3600 * 24 * 30 * 6
	app.Sessions.Store = nanostore.New(db.DB.Collection("Session"))

	// Security
	configureHTTPS(app)

	// Pages
	pages.Configure(app)

	// Middleware
	/*app.Use(
		middleware.Log(),
		middleware.Session(),
	)*/

	// API
	// db.API.Install(app)

	// Development server configuration
	if db.IsDevelopment() {
		app.Config.Domain = "beta.omany.app"

		// Test connectivity
		app.OnStart(testConnectivity)
	}

	// Authentication
	// auth.Install(app)

	// Close the database node on shutdown
	// app.OnEnd(arn.Node.Close)

	// Check that this is the server
	if !db.Node.IsServer() && !db.IsTest() {
		panic("Another program is currently running as the database server")
	}

	// Specify test routes
	for route, examples := range routetests.All() {
		app.Test(route, examples)
	}

	return app
}
