package main

import (
	"httpserver/web"
)

func main() {
	web.AddRoute("/checkstatus")
	web.Run(":37700")
}

