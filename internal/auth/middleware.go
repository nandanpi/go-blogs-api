package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var JWTKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(JWTKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func ProtectedRoute(f func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) == 0 {
			c.JSON(401, gin.H{"message": "Missing Authorization Header"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		_, err := verifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		f(c)
	}
}
