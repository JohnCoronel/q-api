package main

import (
	"github.com/johncoronel/q-api/server"
)

type Env struct {
	
}
func main() {
	server.Server().ListenAndServe()
}
