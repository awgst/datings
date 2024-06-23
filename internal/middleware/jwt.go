package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}
		if !strings.Contains(tokenString, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("signing method is invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, errors.New("signing method is invalid")
			}

			return []byte(cfg.JWT.Secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}

		userID, ok := claims["id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}

		// Set user to context
		c.Set("userID", userID)
		c.Next()
	}
}
