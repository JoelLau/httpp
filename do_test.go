package httpp

import (
	"fmt"
	"net/http"
	"net/url"
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
	c.CallHistory = append(c.CallHistory, fmt.Sprintf("%s %s", req.Method, req.URL))

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
