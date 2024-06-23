package feed

import (
	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/middleware"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

type NewRoutesParams struct {
	R             *gin.RouterGroup
	ErrorLogger   logger.Interface
	Configuration *config.Config
	FeedUsecase   usecase.FeedUsecase
}

func NewRoutes(param NewRoutesParams) {
	handler := NewHandler(param.FeedUsecase, param.ErrorLogger)
	routes := param.R.Group("/feed")
	{
		routes.Use(middleware.JwtAuth(param.Configuration))
		routes.POST("/swipe", handler.Swipe)
		routes.GET("", handler.Recommendation)
	}
}
