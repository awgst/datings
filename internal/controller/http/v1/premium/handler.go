package premium

import (
	"fmt"
	"net/http"

	premiumrequest "github.com/awgst/datings/internal/entity/request/premium"
	"github.com/awgst/datings/internal/usecase"

	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/controller/http/validator"
	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	premiumUsecase usecase.PremiumUsecase
	errorLogger    logger.Interface
}

func NewHandler(premiumUsecase usecase.PremiumUsecase, errorLogger logger.Interface) Handler {
	return Handler{
		premiumUsecase: premiumUsecase,
		errorLogger:    errorLogger,
	}
}

func (h Handler) Purchase(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	var req premiumrequest.CreatePremiumRequest
	errMessages, err := validator.Validate(ctx, &req)
	if err != nil {
		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.premium] Buy: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to buy", err.Error()))
		return
	}

	if len(errMessages) > 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			response.JSON(false, "Failed to buy", customerror.Error{
				Code: customerror.ErrorCodeInvalidRequest,
				Err:  errMessages,
			}),
		)
		return
	}

	if err := h.premiumUsecase.Create(int(userID), req); err != nil {
		customErr, ok := err.(customerror.Error)
		if ok {
			ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to buy", customErr))
			return
		}

		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.premium] Buy: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(true, "Success", nil))
}
