package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"mgo-gin/app/model"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("uit_secret_key")

type Claims struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func GenerateJWTToken(user model.User) string {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		Roles: user.Roles,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err!=nil{
		logrus.Print(err)
	}
	return tokenString
}
