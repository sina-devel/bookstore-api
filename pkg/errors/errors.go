package errors

import "net/http"

type (
	kind uint

	serverError struct {
		kind    kind
		message string
	}
)

const (
	_ kind = iota
	KindInvalid
	KindNotFound
	KindUnauthorized
	KindUnexpected
	KindNotAllowed
)

var (
	httpErrors = map[kind]int{
		KindInvalid:      http.StatusBadRequest,
		KindNotFound:     http.StatusNotFound,
		KindUnauthorized: http.StatusUnauthorized,
		KindUnexpected:   http.StatusInternalServerError,
		KindNotAllowed:   http.StatusMethodNotAllowed,
	}
)

//New is Constructor of the errors package
func New(kind kind, msg string) error {
	return serverError{
		kind:    kind,
		message: msg,
	}
}

//Error return message of error
func (e serverError) Error() string {
	return e.message
}

//HttpError convert kind of error to Http status error
//if error type is not serverError return 400 status code
func HttpError(err error) (string, int) {
	serverErr, ok := err.(serverError)
	if !ok {
		return err.Error(), http.StatusBadRequest
	}

	code, ok := httpErrors[serverErr.kind]
	if !ok {
		return serverErr.message, http.StatusBadRequest
	}

	return serverErr.message, code

}
