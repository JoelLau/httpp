package httpp

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(urlStr string, opts ...OptFunc[DoParam]) (res Response, err error) {
	URL, err := url.Parse(urlStr)
	if err != nil {
		err = fmt.Errorf("invalid url: %w", err)
		return
	}
	req := Request{Method: http.MethodGet, URL: URL}
	return Do(req, opts...)
}

func WithParams(params io.ReadCloser) OptFunc[DoParam] {
	return func(oldParams DoParam) (newParams DoParam, err error) {
		newParams = oldParams
		newParams.Request.Body = params

		return
	}
}
