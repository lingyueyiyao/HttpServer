package main

import (
	"httpserver/web"
	"net/http"
)

func main() {
	web.AddRoute("/checkstatus")
	http.ListenAndServe(":37700", nil)
}

