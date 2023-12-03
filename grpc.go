package statuserr

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func codeFromStatus(status int) codes.Code {
	switch status {
	case http.StatusOK:
		return codes.OK
	case StatusCanceled:
		return codes.Canceled
	case StatusUnknown:
		return codes.Unknown
	case http.StatusUnprocessableEntity:
		return codes.InvalidArgument
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusBadRequest:
		return codes.FailedPrecondition
	case http.StatusPreconditionRequired:
		return codes.Aborted
	case http.StatusRequestedRangeNotSatisfiable:
		return codes.OutOfRange
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusInsufficientStorage:
		return codes.DataLoss
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	default:
		return codes.Unknown
	}
}

func statusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return StatusCanceled
	case codes.Unknown:
		return StatusUnknown
	case codes.InvalidArgument:
		return http.StatusUnprocessableEntity
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusPreconditionRequired
	case codes.OutOfRange:
		return http.StatusRequestedRangeNotSatisfiable
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusBadGateway
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInsufficientStorage
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return StatusUnknown
	}
}

func wrapStatus(err error) error {
	if err == nil {
		return nil
	}
	return status.Error(codeFromStatus(Status(err)), err.Error())
}

func unwrapStatus(err error) error {
	if err == nil {
		return nil
	}
	s, ok := status.FromError(err)
	if !ok {
		return &statusError{StatusUnknown, err}
	}
	code := s.Code()
	if code == codes.OK {
		return nil
	}
	return &statusError{statusFromCode(code), errors.New(s.Message())}
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		return resp, wrapStatus(err)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return wrapStatus(handler(srv, ss))
	}
}

func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return unwrapStatus(invoker(ctx, method, req, reply, cc, opts...))
	}
}

func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		stream, err := streamer(ctx, desc, cc, method, opts...)
		return stream, unwrapStatus(err)
	}
}
