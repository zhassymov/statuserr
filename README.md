# statuserr

---
The **statuserr** package provides for the creation and transmission of errors with a status code between services.

The package completely supports errors wrapping.

```
go get github.com/zhassymov/statuserr
```

| gRPC code                 | http status code                         |
|---------------------------|------------------------------------------|
| 0 OK                      | 200 OK                                   |
| 1 Canceled                | 499 Client Closed Request                |
| 2 Unknown                 | 520 Web Server Returned an Unknown Error |
| 3 InvalidArgument         | 422 Unprocessable Entity                 |
| 4 DeadlineExceeded        | 504 Gateway Timeout                      |
| 5 NotFound                | 404 Not Found                            |
| 6 AlreadyExists           | 409 Conflict                             |
| 7 PermissionDenied        | 403 Forbidden                            |
| 8 ResourceExhausted       | 429 Too Many Requests                    |
| 9 FailedPrecondition      | 400 Bad Request                          |
| 10 Aborted                | 428 Precondition Required                |
| 11 OutOfRange             | 416 Requested Range Not Satisfiable      |
| 12 Unimplemented          | 501 Not Implemented                      |
| 13 Internal               | 502 Bad Gateway                          |
| 14 Unavailable            | 503 Service Unavailable                  |
| 15 DataLoss               | 507 Insufficient Storage                 |
| 16 Unauthenticated        | 401 Unauthorized                         |

## Usage:

### http

Create an error, for example: `statuserr.NotFound(fmt.Errorf("entity with id %d not found", id))`

Use `statuserr.Status(err)` in your handler or middleware to get error http status code.

### gRPC

Use the interceptors on gRPC Server side:
- statuserr.UnaryServerInterceptor()
- statuserr.StreamServerInterceptor()

Use the interceptors on gRPC Client side:
- statuserr.UnaryClientInterceptor()
- statuserr.StreamClientInterceptor()

If you create an error, for example: `statuserr.NotFound(fmt.Errorf("entity with id %d not found", id))`, it is converted with gRPC code "5 NotFound" on the server side and with "404 Not Found" on the client side.
If client is been as a gPRC server for a third service, this error is converted with gRPC code "5 NotFound" again. And so on down the chain.
