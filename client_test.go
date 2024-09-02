package httpp

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
