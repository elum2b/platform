package version

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestLatestCachesPublishedVersion(t *testing.T) {
	var tokenRequests atomic.Int32

	client := versionClient(&tokenRequests, "v1.2.3", time.Hour)

	for range 2 {
		latest, err := client.Latest(context.Background())
		if err != nil {
			t.Fatalf("Latest() error = %v", err)
		}

		if latest != "v1.2.3" {
			t.Errorf("Latest() = %q, want %q", latest, "v1.2.3")
		}
	}

	if requests := tokenRequests.Load(); requests != 1 {
		t.Errorf("token requests = %d, want %d", requests, 1)
	}
}

func TestLatestRefreshesExpiredCache(t *testing.T) {
	var tokenRequests atomic.Int32

	client := versionClient(&tokenRequests, "v1.2.3", 0)

	for range 2 {
		if _, err := client.Latest(context.Background()); err != nil {
			t.Fatalf("Latest() error = %v", err)
		}
	}

	if requests := tokenRequests.Load(); requests != 2 {
		t.Errorf("token requests = %d, want %d", requests, 2)
	}
}

func TestImageVersionUsesOCIAnnotation(t *testing.T) {
	version := imageVersion(map[string]string{
		ociVersionLabel: "v1.2.3",
	})
	if version != "v1.2.3" {
		t.Errorf("imageVersion() = %q, want %q", version, "v1.2.3")
	}
}

func versionClient(
	requests *atomic.Int32,
	version string,
	cacheTTL time.Duration,
) *Client {
	httpClient := &http.Client{Transport: roundTripFunc(func(
		request *http.Request,
	) (*http.Response, error) {
		body := ""

		switch request.URL.Path {
		case "/token":
			requests.Add(1)
			body = `{"token":"token"}`
		case "/v2/elum2b/platform/manifests/latest":
			body = `{"config":{"digest":"sha256:config"}}`
		case "/v2/elum2b/platform/blobs/sha256:config":
			body = `{"config":{"Labels":{"io.elum2b.platform.version":"` +
				version + `"}}}`
		default:
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader("")),
				Header:     make(http.Header),
				Request:    request,
			}, nil
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
			Request:    request,
		}, nil
	})}

	return NewClient(httpClient, "https://ghcr.io", "elum2b/platform", cacheTTL)
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (roundTrip roundTripFunc) RoundTrip(
	request *http.Request,
) (*http.Response, error) {
	return roundTrip(request)
}
