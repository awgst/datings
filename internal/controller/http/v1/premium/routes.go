package premium

import (
	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/middleware"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

type NewRoutesParams struct {
	R              *gin.RouterGroup
	ErrorLogger    logger.Interface
	Configuration  *config.Config
	PremiumUsecase usecase.PremiumUsecase
}

func NewRoutes(param NewRoutesParams) {
	handler := NewHandler(param.PremiumUsecase, param.ErrorLogger)
	routes := param.R.Group("/premium")
	{
		routes.Use(middleware.JwtAuth(param.Configuration))
		routes.POST("", handler.Purchase)
	}
}
