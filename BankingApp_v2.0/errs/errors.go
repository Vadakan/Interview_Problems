package errs

import "net/http"

type AppError struct{
	Code    int     `json:",omitempty"`
	Message string  `json:"message"`
}

func(s AppError) AsMessage() *AppError{

	return &AppError{
		Message: s.Message,
	}
}

func NewNotFoundError(message string) *AppError{
	return &AppError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedDatabaseError(message string) *AppError{
	return &AppError{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}
