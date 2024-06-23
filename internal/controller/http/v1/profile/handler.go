package profile

import (
	"fmt"
	"net/http"

	"github.com/awgst/datings/internal/controller/http/response"
	"github.com/awgst/datings/internal/customerror"
	profilerequest "github.com/awgst/datings/internal/entity/request/profile"
	profileresponse "github.com/awgst/datings/internal/entity/response/profile"
	"github.com/awgst/datings/internal/usecase"
	"github.com/awgst/datings/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase usecase.UserUsecase
	errorLogger logger.Interface
}

func NewHandler(userUsecase usecase.UserUsecase, errorLogger logger.Interface) Handler {
	return Handler{
		userUsecase: userUsecase,
		errorLogger: errorLogger,
	}
}

func (h Handler) Profile(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	user, err := h.userUsecase.FindByID(int(userID))
	if err != nil {
		customErr, ok := err.(customerror.Error)
		if ok {
			ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to get profile", customErr))
			return
		}

		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.profile] Profile: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(true, "Success", new(profileresponse.ProfileResponse).Make(user)))
}

func (h Handler) UpdateProfile(ctx *gin.Context) {
	userID := ctx.GetFloat64("userID")
	if userID == 0 {
		ctx.JSON(http.StatusBadRequest, response.UnauthorizedResponse)
		return
	}

	var req profilerequest.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.profile] UpdateProfile: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to update profile", nil))
		return
	}

	err := h.userUsecase.Update(int(userID), req)
	if err != nil {
		customErr, ok := err.(customerror.Error)
		if ok {
			ctx.JSON(http.StatusBadRequest, response.JSON(false, "Failed to update profile", customErr))
			return
		}

		h.errorLogger.Error(fmt.Errorf("[internal.controller.http.v1.profile] UpdateProfile: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.JSON(false, "Something went wrong", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.JSON(true, "Success", nil))
}
