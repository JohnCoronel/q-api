package main

import (
	"github.com/johncoronel/q-api/server"
)

func main() {
	server.Server().ListenAndServe()
}
