package auth

import (
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.RouterGroup, errorLogger logger.Interface, authUsecase usecase.AuthUsecase) {
	handler := NewHandler(authUsecase, errorLogger)
	routes := r.Group("/auth")
	{
		routes.POST("/signup", handler.SignUp)
		routes.POST("/login", handler.Login)
	}
}
