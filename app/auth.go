package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/LordRahl90/contacts_manager/models"
	"github.com/LordRahl90/contacts_manager/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

//JwtAuthentication middleware
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		notAuth := []string{"/api/user/register", "/api/user/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				fmt.Println(value)
				next.ServeHTTP(w, r)
				return
			}
		} //checks the paths to confirm if the route is not a protected route.

		headerToken := r.Header.Get("Authorization")
		response := make(map[string]interface{})
		if headerToken == "" {
			response = utils.Message(false, "Token Not found")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		splittedToken := strings.Split(headerToken, " ")
		println(splittedToken[0])

		if len(splittedToken) != 2 {
			response = utils.Message(false, "Invalid/Malformed Token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		tokenPart := splittedToken[1] //this contains the token
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = utils.Message(false, "Malformed Authentication Token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		if !token.Valid {
			response = utils.Message(false, "Invalid Token")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		//valid token at this point.
		fmt.Println("Token Valid at this point")
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
