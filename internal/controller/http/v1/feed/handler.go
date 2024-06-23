package feed

import (
	"fmt"
	"net/http"

	"github.com/awgst/datings/internal/entity/model"
	feedrequest "github.com/awgst/datings/internal/entity/request/feed"

	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/controller/http/validator"
	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	feedUsecase usecase.FeedUsecase
	errorLogger logger.Interface
}

func NewHandler(feedUsecase usecase.FeedUsecase, errorLogger logger.Interface) Handler {
	return Handler{
		feedUsecase: feedUsecase,
		errorLogger: errorLogger,
	}
}

func (h Handler) Swipe(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	var req feedrequest.SwipeRequest
	errMessages, err := validator.Validate(ctx, &req)
	if err != nil {
		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.feed] Swipe: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to swipe", err.Error()))
		return
	}

	if len(errMessages) > 0 {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			response.JSON(false, "Failed to swipe", customerror.Error{
				Code: customerror.ErrorCodeInvalidRequest,
				Err:  errMessages,
			}),
		)
		return
	}

	if err := h.feedUsecase.Swipe(model.User{
		ID: int(userID),
	}, req); err != nil {
		customErr, ok := err.(customerror.Error)
		if ok {
			ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to swipe", customErr))
			return
		}

		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.feed] Swipe: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(true, "Success", nil))
}
