package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/entity/model"
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

		userID, ok := claims["user_id"].(float64)
		if !ok {
			fmt.Println("3")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			fmt.Println("2")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.UnauthorizedResponse)
			return
		}

		var premiumFeature string
		premium, exists := claims["premium"]
		if exists {
			premiumFeature, ok = premium.(string)
			if !ok || premiumFeature == "none" {
				premiumFeature = ""
			}
		}

		// Set user to context
		c.Set("userID", userID)
		c.Set("user", model.User{
			ID:    int(userID),
			Email: email,
			Premium: &model.Premium{
				Feature: model.PremiumFeature(premiumFeature),
			},
		})
		c.Next()
	}
}
