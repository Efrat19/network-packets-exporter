package src

import (
	"fmt"
	"time"
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func logAccess() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(time.Now().Format("2006-01-02T15:04:05.000Z"),r.Method,r.URL,r.UserAgent())
			h.ServeHTTP(w, r)
		})
	}
}

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}