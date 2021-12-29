package helpers

import "net/http"

func GenerateErrorCode(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNoContent
	case ErrDbServer:
		return http.StatusInternalServerError
	case ErrBadReq:
		return http.StatusBadGateway
	case ErrForbidden:
		return http.StatusForbidden
	}
	return http.StatusNotFound
}
