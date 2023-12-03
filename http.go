package statuserr

import (
	"errors"
	"net/http"
)

const (
	StatusCanceled = 499 // Client Closed Request
	StatusUnknown  = 520 // Web Server Returned an Unknown Error
)

func Status(err error) int {
	if err == nil {
		return http.StatusOK
	}
	var cause interface{ Status() int }
	if errors.As(err, &cause) {
		return cause.Status()
	}
	return http.StatusInternalServerError
}
