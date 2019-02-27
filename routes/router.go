package routes

import (
	mux "github.com/gorilla/mux"
)

//Router returns a gorilla mux router with routes
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", ping)
	return r
}
