package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("remote address: ", r.RemoteAddr)
		log.Println("forwarded for: ", r.Header.Get("X-Forwarded-For"))
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
