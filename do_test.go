package httpp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type MockHttpClient struct {
	CallHistory []string
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if len(c.CallHistory) == 0 {
		c.CallHistory = make([]string, 0)
	}

	if req == nil {
		req = &http.Request{}
	}

	if req.Body == nil {
		req.Body = io.NopCloser(strings.NewReader(""))
	}

	b, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("failed to read req.body")
	}

	asdf := fmt.Sprintf("%s %s %s", req.Method, req.URL, string(b))
	asdf = strings.Trim(asdf, " ")
	c.CallHistory = append(c.CallHistory, asdf)

	return nil, nil
}

func TestDo(t *testing.T) {
	t.Run("makes http call via client", func(t *testing.T) {
		client := MockHttpClient{}
		URL, _ := url.Parse("https://catfact.ninja/fact")
		req := Request{Method: http.MethodGet, URL: URL}
		_, _ = Do(req, WithClient(&client))

		have := client.CallHistory
		want := []string{"GET https://catfact.ninja/fact"}
		if diff := cmp.Diff(want, have); diff != "" {
			t.Fatalf("mock http request history mismatch (-want +have):\n %s", diff)
		}
	})
}
