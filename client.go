package httpp

import (
	"fmt"
)

func HttpClient(opts ...ClientOptFunc) (c *Client, err error) {
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

type ClientOptFunc OptFunc[Client]
