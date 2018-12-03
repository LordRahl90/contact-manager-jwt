package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

//Logger - Middleware for Logging all the incoming request
var Logger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		fmt.Println("Request got here")
		next.ServeHTTP(w, r)
	})
}
