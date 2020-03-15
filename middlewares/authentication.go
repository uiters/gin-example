package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mgo-gin/app/model"
	"net/http"
	"strings"
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
	expirationTime := time.Now().Add(24 * time.Hour)
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

func RequireAuthenticated()  gin.HandlerFunc {
	return func (c *gin.Context){
		token:=c.GetHeader("Authorization")
		if token ==""{
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		jwtToken :=strings.Split(token,"Bearer ")
		// Initialize a new instance of `Claims`
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(jwtToken[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Abort()
				c.Writer.WriteHeader(http.StatusUnauthorized)
			}
			c.Abort()
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
		if !tkn.Valid {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
		}
		return
	}
}
