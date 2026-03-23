package main

import (
	"fmt"
	"net/http"
)

// declare the handlere thats returns plain text response with
// application status, env and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: Available")
	fmt.Fprintf(w, "environment: %s \n ", app.config.env)
	fmt.Fprint(w, "version: %s \n", version)

}
