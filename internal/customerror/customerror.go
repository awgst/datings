package customerror

import "fmt"

type ErrorCode string

const (
	ErrorCodeInvalidCredentials ErrorCode = "invalid_credentials"
	ErrorCodeNotFound           ErrorCode = "not_found"
	ErrorCodeAlreadyExists      ErrorCode = "already_exists"
	ErrorCodeInvalidRequest     ErrorCode = "invalid_request"
)

var (
	ErrorNotFound = Error{
		Code: ErrorCodeNotFound,
		Err:  "not found",
	}
)

type Error struct {
	Code ErrorCode   `json:"code"`
	Err  interface{} `json:"errors"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Code, e.Err)
}
