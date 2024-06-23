package auth

import (
	"fmt"
	"net/http"

	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/controller/http/validator"
	"github.com/awgst/datings/internal/customerror"
	authrequest "github.com/awgst/datings/internal/entity/request/auth"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authUsecase usecase.AuthUsecase
	errorLogger logger.Interface
}

func NewHandler(authUsecase usecase.AuthUsecase, errorLogger logger.Interface) *Handler {
	return &Handler{
		authUsecase: authUsecase,
		errorLogger: errorLogger,
	}
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var req authrequest.SignupRequest
	errMessages, err := validator.Validate(ctx, &req)
	if err != nil {
		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.auth] Signup: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to sign up", err.Error()))
		return
	}

	if len(errMessages) > 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			response.JSON(false, "Failed to sign up", errMessages),
		)
		return
	}

	result, err := h.authUsecase.SignUp(req)
	if err != nil {
		customErr, ok := err.(customerror.Error)
		if ok {
			ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to sign up", customErr))
			return
		}

		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.auth] Signup: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(true, "Success", result))
}
