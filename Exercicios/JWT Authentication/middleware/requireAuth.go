package middleware

import (
	"jwt-authentication/initializers"
	"jwt-authentication/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie from the request
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Parse the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > token.Claims.(jwt.MapClaims)["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token has expired",
		})
		c.Abort()
		return
	}

	//Find the user with token id
	var user models.User
	initializers.DB.First(&user, "id = ?", token.Claims.(jwt.MapClaims)["id"])
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	// Set the user in the context
	c.Set("user", user)
	c.Next()
}
