package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jamestjw/coup-vin/models"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type AccessTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateTokenPair(u models.User) (map[string]string, error) {
	// Access Token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &AccessTokenClaims{
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			Subject:   strconv.Itoa(u.ID),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   strconv.Itoa(u.ID),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}
