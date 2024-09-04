package httpp

import (
	"io"
	"strings"
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

func TestPost(t *testing.T) {
	t.Run("no args other than mock", func(t *testing.T) {
		client := &MockHttpClient{}
		_, err := Post("https://catfact.ninja/fact",
			HttpBody{ContentType: "application/json",
				Body: io.NopCloser(strings.NewReader(`{"fact":"70% of your cat's life is spent asleep.","length":39}`))},
			WithClient(client))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		have := client.CallHistory
		want := []string{`POST https://catfact.ninja/fact {"fact":"70% of your cat's life is spent asleep.","length":39}`}
		if diff := cmp.Diff(want, have); diff != "" {
			t.Fatalf("mock http request history mismatch (-want +have):\n %s", diff)
		}
	})
}
