package statuserr

import "net/http"

type statusError struct {
	status int
	err    error
}

// Unwrap provides compatibility for Go 1.13 error chains
func (e *statusError) Unwrap() error { return e.err }

// Cause provides compatibility for github.com/pkg/errors error chains
func (e *statusError) Cause() error { return e.err }

func (e *statusError) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *statusError) Status() int {
	return e.status
}

////////////////////////////////
//     client-side errors     //
////////////////////////////////s

func BadRequest(err error) error {
	return &statusError{http.StatusBadRequest, err}
}

func Unauthorized(err error) error {
	return &statusError{http.StatusUnauthorized, err}
}

func PaymentRequired(err error) error {
	return &statusError{http.StatusPaymentRequired, err}
}

func Forbidden(err error) error {
	return &statusError{http.StatusForbidden, err}
}

func NotFound(err error) error {
	return &statusError{http.StatusNotFound, err}
}

func MethodNotAllowed(err error) error {
	return &statusError{http.StatusMethodNotAllowed, err}
}

func NotAcceptable(err error) error {
	return &statusError{http.StatusNotAcceptable, err}
}

func ProxyAuthRequired(err error) error {
	return &statusError{http.StatusProxyAuthRequired, err}
}

func RequestTimeout(err error) error {
	return &statusError{http.StatusRequestTimeout, err}
}

func Conflict(err error) error {
	return &statusError{http.StatusConflict, err}
}

func Gone(err error) error {
	return &statusError{http.StatusGone, err}
}

func LengthRequired(err error) error {
	return &statusError{http.StatusLengthRequired, err}
}

func PreconditionFailed(err error) error {
	return &statusError{http.StatusPreconditionFailed, err}
}

func RequestEntityTooLarge(err error) error {
	return &statusError{http.StatusRequestEntityTooLarge, err}
}

func RequestURITooLong(err error) error {
	return &statusError{http.StatusRequestURITooLong, err}
}

func UnsupportedMediaType(err error) error {
	return &statusError{http.StatusUnsupportedMediaType, err}
}

func RequestedRangeNotSatisfiable(err error) error {
	return &statusError{http.StatusRequestedRangeNotSatisfiable, err}
}

func ExpectationFailed(err error) error {
	return &statusError{http.StatusExpectationFailed, err}
}

func MisdirectedRequest(err error) error {
	return &statusError{http.StatusMisdirectedRequest, err}
}

func UnprocessableEntity(err error) error {
	return &statusError{http.StatusUnprocessableEntity, err}
}

func Locked(err error) error {
	return &statusError{http.StatusLocked, err}
}

func FailedDependency(err error) error {
	return &statusError{http.StatusFailedDependency, err}
}

func TooEarly(err error) error {
	return &statusError{http.StatusTooEarly, err}
}

func UpgradeRequired(err error) error {
	return &statusError{http.StatusUpgradeRequired, err}
}

func PreconditionRequired(err error) error {
	return &statusError{http.StatusPreconditionRequired, err}
}

func TooManyRequests(err error) error {
	return &statusError{http.StatusTooManyRequests, err}
}

func RequestHeaderFieldsTooLarge(err error) error {
	return &statusError{http.StatusRequestHeaderFieldsTooLarge, err}
}

func UnavailableForLegalReasons(err error) error {
	return &statusError{http.StatusUnavailableForLegalReasons, err}
}

func Canceled(err error) error {
	return &statusError{StatusCanceled, err}
}

////////////////////////////////
//     server-side errors     //
////////////////////////////////

func InternalServerError(err error) error {
	return &statusError{http.StatusInternalServerError, err}
}

func NotImplemented(err error) error {
	return &statusError{http.StatusNotImplemented, err}
}

func BadGateway(err error) error {
	return &statusError{http.StatusBadGateway, err}
}

func ServiceUnavailable(err error) error {
	return &statusError{http.StatusServiceUnavailable, err}
}

func GatewayTimeout(err error) error {
	return &statusError{http.StatusGatewayTimeout, err}
}

func HTTPVersionNotSupported(err error) error {
	return &statusError{http.StatusHTTPVersionNotSupported, err}
}

func VariantAlsoNegotiates(err error) error {
	return &statusError{http.StatusVariantAlsoNegotiates, err}
}

func InsufficientStorage(err error) error {
	return &statusError{http.StatusInsufficientStorage, err}
}

func LoopDetected(err error) error {
	return &statusError{http.StatusLoopDetected, err}
}

func NotExtended(err error) error {
	return &statusError{http.StatusNotExtended, err}
}

func NetworkAuthenticationRequired(err error) error {
	return &statusError{http.StatusNetworkAuthenticationRequired, err}
}

func Unknown(err error) error {
	return &statusError{StatusUnknown, err}
}
