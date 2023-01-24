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
	ErrInternalServer = NewApiError("internal server error")
)
