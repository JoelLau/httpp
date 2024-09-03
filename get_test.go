package httpp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGet(t *testing.T) {
	t.Run("no args other than mock", func(t *testing.T) {
		client := &MockHttpClient{}
		_, err := Get("https://catfact.ninja/fact", WithClient(client))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		have := client.CallHistory
		want := []string{"GET https://catfact.ninja/fact"}
		if diff := cmp.Diff(want, have); diff != "" {
			t.Fatalf("mock http request history mismatch (-want +have):\n %s", diff)
		}
	})
}
