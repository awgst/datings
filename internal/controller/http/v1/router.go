// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/controller/http/v1/auth"
	"github.com/awgst/datings/internal/usecase"
)

// NewRouter -.
func NewRouter(handler *gin.Engine, uc *usecase.Usecase) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.CustomRecovery(func(ctx *gin.Context, err interface{}) {
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
	}))

	handler.GET("/healthz", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, response.JSON(true, "Ok", nil)) })

	v1 := handler.Group("/v1")
	{
		auth.NewRoutes(v1, uc.App.Logger, uc.Auth)
	}
}
