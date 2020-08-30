package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamestjw/coup-vin/models"
)

// Middleware handles the authentication of a user by checking the JWT token in
// the Authorization header
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtKey), nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

				id, err := strconv.Atoi(claims["sub"].(string))
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid Token"))
					fmt.Print((err))
					return
				}

				user := models.FindUserByID(id)

				if user == nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid Token"))
					fmt.Print((err))
					return
				}

				ctx := context.WithValue(r.Context(), "user", user)
				// Access user values in handlers like this
				// user, _ := r.Context().Value("user").(models.User)

				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func CurrentUser(r *http.Request) (models.User, error) {
	user, ok := r.Context().Value("user").(models.User)
	if !ok {
		return models.User{}, fmt.Errorf("current user is not found")
	}
	return user, nil
}
