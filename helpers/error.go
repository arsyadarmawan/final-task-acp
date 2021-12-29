package helpers

import "errors"

var (
	ErrNotFound      = errors.New("data not found")
	ErrForbidden     = errors.New("Forbidden")
	ErrDbServer      = errors.New("db server error")
	ErrBadReq        = errors.New("Bad Req")
	ErrEmptyUrlParam = errors.New("Empty url param")
)
