package httpp

import (
	"fmt"
	"net/http"
)

// makes request using the http client (default this library, not net/http's)
func Do(req Request, opts ...DoOptFunc) (response Response, err error) {
	client, err := HttpClient()
	if err != nil {
		err = fmt.Errorf("error creating base http client: %w", err)
		return
	}

	param := DoParam{Client: client, Request: req}
	for _, opt := range opts {
		param, err = opt(param)
		if err != nil {
			err = fmt.Errorf("error applying Do() opt func: %w", err)
			return
		}
	}

	_, err = param.Client.Do(&param.Request)
	if err != nil {
		err = fmt.Errorf("error while making http request: %w", err)
		return
	}

	return
}

type DoParam struct {
	Client  HttpDoer
	Request Request
}

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type DoOptFunc = OptFunc[DoParam]

func WithClient(client HttpDoer) DoOptFunc {
	return func(old DoParam) (param DoParam, err error) {
		param = old
		param.Client = client

		return
	}
}
