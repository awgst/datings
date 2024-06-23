package auth

import (
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.RouterGroup, errorLogger logger.Interface, userUsecase usecase.AuthUsecase) {
	handler := NewHandler(userUsecase, errorLogger)
	routes := r.Group("/auth")
	{
		routes.POST("/signup", handler.SignUp)
	}
}
