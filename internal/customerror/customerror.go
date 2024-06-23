package customerror

import "fmt"

type ErrorCode string

const (
	ErrorCodeNotFound      ErrorCode = "not_found"
	ErrorCodeAlreadyExists ErrorCode = "already_exists"
)

var (
	ErrorNotFound = Error{
		Code: ErrorCodeNotFound,
		Err:  "not found",
	}
)

type Error struct {
	Code ErrorCode
	Err  interface{}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Code, e.Err)
}
