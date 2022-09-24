package main

import "httpserver/http"

func main() {
	engine := http.NewEngine()
	engine.AddRoute("/hello", func(ctx *http.Context){
		ctx.Writer.Write([]byte("hello"))
	})
	engine.ListenAndServe(":8080")
}

