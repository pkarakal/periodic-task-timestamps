package server

import (
	"fmt"
	"net/http"
)

type ResponseError struct {
	Status      int    `json:"status"`
	Description string `json:"desc"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%s", e.Description)
}

var (
	ErrInvalidT1Timestamp = ResponseError{
		Status:      http.StatusBadRequest,
		Description: fmt.Sprintf("t1 is not a valid timestamp"),
	}
	ErrInvalidT2Timestamp = ResponseError{
		Status:      http.StatusBadRequest,
		Description: fmt.Sprintf("t2 is not a valid timestamp"),
	}
	ErrUnsupportedPeriod = ResponseError{
		Status:      http.StatusConflict,
		Description: "Unsupported period"}
	ErrInvalidTimezone = ResponseError{
		Status:      http.StatusBadRequest,
		Description: fmt.Sprintf("Timezone is invalid"),
	}
	ErrStartDateGreater = ResponseError{
		Status:      http.StatusBadRequest,
		Description: fmt.Sprintf("t1 cannot be greater than t2"),
	}
)
