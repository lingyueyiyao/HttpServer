package http

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: r,
	}
}