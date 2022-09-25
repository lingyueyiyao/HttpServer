package http

import (
	"errors"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *Engine) AddRoute(url string, handler HandlerFunc) error {
	if _, ok := e.router[url]; ok {
		return errors.New("route is exit already")
	}

	e.router[url] = handler
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if handler, ok := e.router[url]; ok {
		handler(NewContext(w, r))
	} else {
		w.WriteHeader(404)
	}
}

func (e *Engine) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, e)
}