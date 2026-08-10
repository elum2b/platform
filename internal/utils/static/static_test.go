package static

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/andybalholm/brotli"
	"github.com/gofiber/fiber/v3"
)

func TestServerServesEmbeddedFiles(t *testing.T) {
	app := newTestApp(t)

	response := request(t, app, http.MethodGet, "/dashboard/assets/app.js", "")
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}

	if contentType := response.Header.Get(
		fiber.HeaderContentType,
	); !strings.HasPrefix(
		contentType,
		"text/javascript",
	) {
		t.Fatalf("content type = %q, want JavaScript", contentType)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if string(body) != "console.log('dashboard')" {
		t.Fatalf("body = %q", body)
	}
}

func TestServerRedirectsDashboardAliases(t *testing.T) {
	app := newTestApp(t)

	for _, requestPath := range []string{
		"/dashboard/",
		"/dashboard/index.html",
	} {
		response := request(t, app, http.MethodGet, requestPath, "")
		response.Body.Close()

		if response.StatusCode != http.StatusPermanentRedirect {
			t.Fatalf(
				"%s status = %d, want %d",
				requestPath,
				response.StatusCode,
				http.StatusPermanentRedirect,
			)
		}

		if location := response.Header.Get(
			fiber.HeaderLocation,
		); location != "/dashboard" {
			t.Fatalf("%s location = %q, want /dashboard", requestPath, location)
		}
	}
}

func TestServerUsesIndexForSolidRouterRoutes(t *testing.T) {
	app := newTestApp(t)

	response := request(
		t,
		app,
		http.MethodGet,
		"/dashboard/workspaces/current",
		"",
	)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if string(body) != "<main>dashboard</main>" {
		t.Fatalf("body = %q", body)
	}
}

func TestServerUsesPrecompressedContent(t *testing.T) {
	app := newTestApp(t)

	tests := []struct {
		name     string
		accept   string
		encoding string
		reader   func(io.Reader) io.Reader
	}{
		{
			name:     "brotli",
			accept:   "gzip, br",
			encoding: "br",
			reader: func(source io.Reader) io.Reader {
				return brotli.NewReader(source)
			},
		},
		{
			name:     "gzip",
			accept:   "gzip",
			encoding: "gzip",
			reader: func(source io.Reader) io.Reader {
				reader, err := gzip.NewReader(source)
				if err != nil {
					t.Fatalf("open gzip: %v", err)
				}

				return reader
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := request(
				t,
				app,
				http.MethodGet,
				"/dashboard",
				test.accept,
			)
			defer response.Body.Close()

			if encoding := response.Header.Get(
				fiber.HeaderContentEncoding,
			); encoding != test.encoding {
				t.Fatalf("encoding = %q, want %q", encoding, test.encoding)
			}

			body, err := io.ReadAll(test.reader(response.Body))
			if err != nil {
				t.Fatalf("read body: %v", err)
			}

			if string(body) != "<main>dashboard</main>" {
				t.Fatalf("body = %q", body)
			}
		})
	}
}

func TestServerHonorsEncodingQuality(t *testing.T) {
	app := newTestApp(t)

	response := request(
		t,
		app,
		http.MethodGet,
		"/dashboard",
		"br;q=0.1, gzip;q=1",
	)
	defer response.Body.Close()

	if encoding := response.Header.Get(
		fiber.HeaderContentEncoding,
	); encoding != "gzip" {
		t.Fatalf("content encoding = %q, want gzip", encoding)
	}
}

func TestServerRejectsUnavailableEncoding(t *testing.T) {
	app := newTestApp(t)

	response := request(
		t,
		app,
		http.MethodGet,
		"/dashboard/assets/video.mp4",
		"identity;q=0, br;q=0, gzip;q=0",
	)
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotAcceptable {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusNotAcceptable,
		)
	}
}

func TestContentTypes(t *testing.T) {
	tests := map[string]struct {
		content  []byte
		expected string
	}{
		"image.jpg": {
			content:  []byte{0xff, 0xd8, 0xff, 0xe0},
			expected: "image/jpeg",
		},
		"image.unknown": {
			content: []byte{
				0x89, 0x50, 0x4e, 0x47,
				0x0d, 0x0a, 0x1a, 0x0a,
			},
			expected: "image/png",
		},
		"module.unknown": {
			content:  []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00},
			expected: "application/wasm",
		},
		"movie.unknown": {
			content: []byte{
				0x00, 0x00, 0x00, 0x18,
				0x66, 0x74, 0x79, 0x70,
				0x6d, 0x70, 0x34, 0x32,
				0x00, 0x00, 0x00, 0x00,
				0x6d, 0x70, 0x34, 0x32,
				0x69, 0x73, 0x6f, 0x6d,
			},
			expected: "video/mp4",
		},
		"script.mjs": {
			content:  []byte("export default {}"),
			expected: "text/javascript; charset=utf-8",
		},
		"style.css": {
			content:  []byte("body {}"),
			expected: "text/css; charset=utf-8",
		},
	}

	for filename, test := range tests {
		t.Run(filename, func(t *testing.T) {
			if actual := contentType(
				filename,
				test.content,
			); actual != test.expected {
				t.Fatalf("content type = %q, want %q", actual, test.expected)
			}
		})
	}
}

func TestServerSupportsETag(t *testing.T) {
	app := newTestApp(t)

	first := request(t, app, http.MethodGet, "/dashboard", "")
	etag := first.Header.Get(fiber.HeaderETag)
	first.Body.Close()

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/dashboard",
		http.NoBody,
	)
	request.Header.Set(fiber.HeaderIfNoneMatch, etag)

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotModified {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusNotModified,
		)
	}
}

func TestServerSupportsByteRanges(t *testing.T) {
	app := newTestApp(t)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/dashboard/assets/video.mp4",
		http.NoBody,
	)
	request.Header.Set(fiber.HeaderAcceptEncoding, "br, gzip")
	request.Header.Set(fiber.HeaderRange, "bytes=2-5")

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusPartialContent {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusPartialContent,
		)
	}

	if contentType := response.Header.Get(
		fiber.HeaderContentType,
	); !strings.HasPrefix(
		contentType,
		"video/mp4",
	) {
		t.Fatalf("content type = %q, want video/mp4", contentType)
	}

	if encoding := response.Header.Get(
		fiber.HeaderContentEncoding,
	); encoding != "" {
		t.Fatalf("content encoding = %q, want identity", encoding)
	}

	if ranges := response.Header.Get(
		fiber.HeaderAcceptRanges,
	); ranges != "bytes" {
		t.Fatalf("accept ranges = %q, want bytes", ranges)
	}

	if contentRange := response.Header.Get(
		fiber.HeaderContentRange,
	); contentRange != "bytes 2-5/10" {
		t.Fatalf("content range = %q, want bytes 2-5/10", contentRange)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if string(body) != "2345" {
		t.Fatalf("body = %q, want 2345", body)
	}
}

func TestServerDoesNotContentEncodeMedia(t *testing.T) {
	app := newTestApp(t)

	response := request(
		t,
		app,
		http.MethodGet,
		"/dashboard/assets/video.mp4",
		"br, gzip",
	)
	defer response.Body.Close()

	if encoding := response.Header.Get(
		fiber.HeaderContentEncoding,
	); encoding != "" {
		t.Fatalf("content encoding = %q, want identity", encoding)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if string(body) != "0123456789" {
		t.Fatalf("body = %q, want complete media file", body)
	}
}

func TestServerSupportsHeadByteRanges(t *testing.T) {
	app := newTestApp(t)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodHead,
		"/dashboard/assets/video.mp4",
		http.NoBody,
	)
	request.Header.Set(fiber.HeaderRange, "bytes=-3")

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusPartialContent {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusPartialContent,
		)
	}

	if contentLength := response.Header.Get(
		fiber.HeaderContentLength,
	); contentLength != "3" {
		t.Fatalf("content length = %q, want 3", contentLength)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if len(body) != 0 {
		t.Fatalf("body length = %d, want 0", len(body))
	}
}

func TestServerSupportsMultipleByteRanges(t *testing.T) {
	app := newTestApp(t)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/dashboard/assets/video.mp4",
		http.NoBody,
	)
	request.Header.Set(fiber.HeaderRange, "bytes=0-1,8-9")

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusPartialContent {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusPartialContent,
		)
	}

	if contentType := response.Header.Get(
		fiber.HeaderContentType,
	); !strings.HasPrefix(
		contentType,
		"multipart/byteranges; boundary=platform-",
	) {
		t.Fatalf("content type = %q, want multipart ranges", contentType)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if !strings.Contains(
		string(body),
		"Content-Range: bytes 0-1/10\r\n\r\n01",
	) {
		t.Fatalf("first range is missing from %q", body)
	}

	if !strings.Contains(
		string(body),
		"Content-Range: bytes 8-9/10\r\n\r\n89",
	) {
		t.Fatalf("second range is missing from %q", body)
	}
}

func TestServerRejectsUnsatisfiableByteRange(t *testing.T) {
	app := newTestApp(t)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/dashboard/assets/video.mp4",
		http.NoBody,
	)
	request.Header.Set(fiber.HeaderRange, "bytes=100-200")

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusRequestedRangeNotSatisfiable {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusRequestedRangeNotSatisfiable,
		)
	}

	if contentRange := response.Header.Get(
		fiber.HeaderContentRange,
	); contentRange != "bytes */10" {
		t.Fatalf("content range = %q, want bytes */10", contentRange)
	}
}

func TestServerHonorsIfRange(t *testing.T) {
	app := newTestApp(t)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/dashboard/assets/app.js",
		http.NoBody,
	)
	request.Header.Set(fiber.HeaderAcceptEncoding, "gzip")
	request.Header.Set(fiber.HeaderIfRange, `"different"`)
	request.Header.Set(fiber.HeaderRange, "bytes=2-5")

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}

	if encoding := response.Header.Get(
		fiber.HeaderContentEncoding,
	); encoding != "gzip" {
		t.Fatalf("content encoding = %q, want gzip", encoding)
	}
}

func newTestApp(t *testing.T) *fiber.App {
	t.Helper()

	server, err := New(fstest.MapFS{
		"index.html": {
			Data: []byte("<main>dashboard</main>"),
		},
		"assets/app.js": {
			Data: []byte("console.log('dashboard')"),
		},
		"assets/video.mp4": {
			Data: []byte("0123456789"),
		},
	})
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	app := fiber.New()
	if err := server.Register(app, "/dashboard"); err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	return app
}

func request(
	t *testing.T,
	app *fiber.App,
	method string,
	path string,
	acceptEncoding string,
) *http.Response {
	t.Helper()

	request := httptest.NewRequestWithContext(
		context.Background(),
		method,
		path,
		http.NoBody,
	)
	if acceptEncoding != "" {
		request.Header.Set(fiber.HeaderAcceptEncoding, acceptEncoding)
	}

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	return response
}
