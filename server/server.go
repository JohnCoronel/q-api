package server

import (
	"net/http"

	r "github.com/johncoronel/q-api/routes"
)

func Server() *http.Server {
	s := &http.Server{
		Addr:    ":8080",
		Handler: r.Router(),
	}
	return s
}
