package feed

import (
	"fmt"
	"net/http"

	"github.com/awgst/datings/internal/entity/model"
	feedquery "github.com/awgst/datings/internal/entity/query/feed"
	feedrequest "github.com/awgst/datings/internal/entity/request/feed"
	feedresponse "github.com/awgst/datings/internal/entity/response/feed"

	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/controller/http/validator"
	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/awgst/datings/pkg/pagination"
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
	userFromCtx, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	user, ok := userFromCtx.(model.User)
	if !ok {
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

	if err := h.feedUsecase.Swipe(user, req); err != nil {
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

func (h Handler) Recommendation(ctx *gin.Context) {
	userFromCtx, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	user, ok := userFromCtx.(model.User)
	if !ok {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	var query feedquery.ListPaginatedQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to get feeds", err.Error()))
		return
	}

	if query.Page == 0 || !user.HasUnlimitedSwipe() {
		query.Page = 1
	}

	paging := new(pagination.Paginator)
	paging.New(10, query.Page, "/v1/feed")

	profiles, err := h.feedUsecase.Recommendation(user, paging)
	if err != nil {
		customErr, ok := err.(customerror.Error)
		if ok {
			ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to get feeds", customErr))
			return
		}

		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.feed] Recommendation: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.JSON(
			true,
			"Success",
			new(feedresponse.RecommendationResponse).Makes(profiles),
			paging.GetLinks(),
		),
	)
}
