package framework

import (
    "net/http"
)

type Router struct {
    mux *http.ServeMux
}

func NewRouter() *Router {
    return &Router{mux: http.NewServeMux()}
}

func (r *Router) Handle(pattern string, handler http.HandlerFunc) {
    r.mux.HandleFunc(pattern, handler)
}

func (r *Router) ServeStatic(prefix, dir string) {
    r.mux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(dir))))
}

func (r *Router) Serve() http.Handler {
    return r.mux
}