package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequestValidator interface {
	GetJsonFieldName(string) string
	ErrMessages() map[string]map[string]string
}

func Validate(ctx *gin.Context, req RequestValidator) (map[string]string, error) {
	var errMessages = map[string]string{}
	if err := ctx.ShouldBind(req); err != nil {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return map[string]string{}, err
		}

		for _, e := range validationErr {
			fieldJSONName := req.GetJsonFieldName(e.Field())
			errMessages[fieldJSONName] = req.ErrMessages()[fieldJSONName][e.ActualTag()]
		}
	}

	return errMessages, nil
}
