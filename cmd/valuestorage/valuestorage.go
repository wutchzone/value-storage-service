package main

import (
	"net/http"
)

func main() {
	r := InitRoutes()

	http.ListenAndServe(":7081", r)
}
