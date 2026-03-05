package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /getByName/{name}", app.getEmployeeByName)
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("GET /getById/{id}", app.getEmployeeById)
	return mux
}
