package middleware

import (
	"fmt"
	"main/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ValidToken ...
func ValidToken(c *gin.Context) (*jwt.Token, bool) {
	// Get token from cookie and check if valid
	tokenString, err := c.Cookie("token")
	if err != nil {
		return nil, false
	}
	token, err := service.JWTAuthService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return token, true
	}
	fmt.Println(err)
	return nil, false

}
