package httpp

import "net/http"

type Client = http.Client
type Request = http.Request
type Response = http.Response

// functional way of adding/modifying optional fields in structs
type OptFunc[T any] func(T) (T, error)
