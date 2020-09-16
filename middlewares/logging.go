package middlewares

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		requestDump, err := httputil.DumpRequest(r, true)
		if err == nil {
			log.Println(string(requestDump))
		}

		next.ServeHTTP(w, r)
	})
}
