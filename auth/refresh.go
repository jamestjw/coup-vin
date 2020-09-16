package auth

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamestjw/coup-vin/models"
)

func RefreshToken(refreshToken string, db models.Datastore) (map[string]string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what we expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("refresh_token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims["sub"].(string))
		if err != nil {
			return nil, err
		}

		user, err := db.FindUserByID(uint(id))

		if err != nil {
			return nil, fmt.Errorf("refresh_token is invalid")
		}
		newTokenPair, err := GenerateTokenPair(user)

		if err != nil {
			return nil, err
		}
		return newTokenPair, err
	}

	return nil, fmt.Errorf("refresh_token is invalid")
}
