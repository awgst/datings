package profile

import (
	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/middleware"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.RouterGroup, errorLogger logger.Interface, cfg *config.Config, userUsecase usecase.UserUsecase) {
	handler := NewHandler(userUsecase, errorLogger)
	routes := r.Group("/profile")
	{
		routes.Use(middleware.JwtAuth(cfg))
		routes.GET("", handler.Profile)
	}
}
