package handlers

type ApiError struct {
	Message string `json:"error"`
}

func NewApiError(msg string) *ApiError {
	return &ApiError{Message: msg}
}

func (e *ApiError) Error() string {
	return e.Message
}

var (
	ErrInternalServer     = NewApiError("500 internal server error")
	ErrUnauthrized        = NewApiError("401 Unauthorized")
	ErrBadRequest         = NewApiError("402 bad request")
	ErrForbidden          = NewApiError("403 forbidden")
	ErrNotFound           = NewApiError("404 not found")
	ErrInvalidCredentials = NewApiError("invalid password or phone number")

	ErrProductNotFound = NewApiError("product not found")
)
