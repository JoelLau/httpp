package httpp

import (
	"fmt"
	"net/http"
)

type Client = http.Client
type Request = http.Request
type Response = http.Response

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func HttpClient(opts ...OptFunc[Client]) (c *Client, err error) {
	client := Client{}

	for _, opt := range opts {
		client, err = opt(client)
		if err != nil {
			err = fmt.Errorf("error applying http client opt func: %w", err)
			return
		}
	}

	return
}

// functional way of adding/modifying optional fields in structs
type OptFunc[T any] func(T) (T, error)

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

type DoOptFunc = OptFunc[DoParam]

func WithClient(client HttpDoer) DoOptFunc {
	return func(old DoParam) (param DoParam, err error) {
		param = old
		param.Client = client

		return
	}
}

func WithRequest(req Request) DoOptFunc {
	return func(old DoParam) (param DoParam, err error) {
		param = old
		param.Request = req

		return
	}
}

type MockHttpClient struct {
	CallHistory []string
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if len(c.CallHistory) == 0 {
		c.CallHistory = make([]string, 0)
	}
	c.CallHistory = append(c.CallHistory, fmt.Sprintf("%s %s", req.Method, req.URL))

	return nil, nil
}
