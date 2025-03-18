package framework

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct {
	Logging func(http.HandlerFunc) http.HandlerFunc
}

func NewMiddleware() Middleware {
	return Middleware{
		Logging: func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				log.Printf("%s %s", r.Method, r.URL.Path)
				next(w, r)
				log.Printf("Completed in %v", time.Since(start))
			}
		},
	}
}
