package main

import "net/http"

func (app *Application) routes () *http.ServeMux{
    mux := http.NewServeMux()
    mux.HandleFunc("/api/healthcheck", app.healthcheck)
    mux.HandleFunc("/api", app.getCreateUser)
    mux.HandleFunc("/api/", app.UserGetUpdateDelete)
    return mux
}
