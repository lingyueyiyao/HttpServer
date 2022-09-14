package web

import "net/http"

func AddRoute(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {})
}

func Run(addr string) {
	http.ListenAndServe(addr, nil)
}
