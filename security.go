package main

import (
	"./db"
	"os"
	"path"

	"github.com/aerogo/aero"
	"github.com/fatih/color"
)

func configureHTTPS(app *aero.Application) {
	fullCertPath := path.Join(db.Root, "security", "fullchain.pem")
	fullKeyPath := path.Join(db.Root, "security", "privkey.pem")

	if _, err := os.Stat(fullCertPath); os.IsNotExist(err) {
		defaultCertPath := path.Join(db.Root, "security", "default", "server.crt")
		err := os.Link(defaultCertPath, fullCertPath)

		if err != nil {
			// Do not panic here, multiple tests could be running this in parallel.
			// Therefore, races can occur (which test writes the link first).
			// Simply log the error and continue as the file should be present.
			color.Red(err.Error())
		}
	}

	if _, err := os.Stat(fullKeyPath); os.IsNotExist(err) {
		defaultKeyPath := path.Join(db.Root, "security", "default", "server.key")
		err := os.Link(defaultKeyPath, fullKeyPath)

		if err != nil {
			// Do not panic here, multiple tests could be running this in parallel.
			// Therefore, races can occur (which test writes the link first).
			// Simply log the error and continue as the file should be present.
			color.Red(err.Error())
		}
	}

	// HTTPS
	app.Security.Load(fullCertPath, fullKeyPath)
}
