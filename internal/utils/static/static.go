// Package static serves embedded static files from memory.
package static

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v3"
)

const indexFile = "index.html"

type file struct {
	compressible bool
	contentType  string
	identity     []byte
	gzip         []byte
	brotli       []byte
	identityETag string
	gzipETag     string
	brotliETag   string
}

type representation struct {
	content  []byte
	encoding string
	etag     string
}

// Server stores all embedded files and their compressed variants in memory.
type Server struct {
	files map[string]file
}

// New reads and compresses every file from files before the HTTP server starts.
func New(files fs.FS) (*Server, error) {
	server := &Server{files: make(map[string]file)}

	if err := fs.WalkDir(files, ".", func(
		filename string,
		entry fs.DirEntry,
		err error,
	) error {
		return server.load(files, filename, entry, err)
	}); err != nil {
		return nil, fmt.Errorf("load static files: %w", err)
	}

	if _, ok := server.files[indexFile]; !ok {
		return nil, fmt.Errorf("static files must contain %s", indexFile)
	}

	return server, nil
}

func (s *Server) load(
	files fs.FS,
	filename string,
	entry fs.DirEntry,
	err error,
) error {
	if err != nil {
		return err
	}

	if entry.IsDir() {
		return nil
	}

	content, err := fs.ReadFile(files, filename)
	if err != nil {
		return fmt.Errorf("read %s: %w", filename, err)
	}

	gzipContent, err := gzipCompress(content)
	if err != nil {
		return fmt.Errorf("gzip %s: %w", filename, err)
	}

	brotliContent, err := brotliCompress(content)
	if err != nil {
		return fmt.Errorf("brotli %s: %w", filename, err)
	}

	fileContentType := contentType(filename, content)

	s.files[filename] = file{
		compressible: isCompressible(fileContentType),
		contentType:  fileContentType,
		identity:     content,
		gzip:         gzipContent,
		brotli:       brotliContent,
		identityETag: etag(content),
		gzipETag:     etag(gzipContent),
		brotliETag:   etag(brotliContent),
	}

	return nil
}

func gzipCompress(content []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := writer.Write(content); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func brotliCompress(content []byte) ([]byte, error) {
	var buffer bytes.Buffer

	writer := brotli.NewWriterLevel(&buffer, brotli.BestCompression)

	if _, err := writer.Write(content); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func contentType(filename string, content []byte) string {
	extension := strings.ToLower(path.Ext(filename))
	contentType := mime.TypeByExtension(extension)

	if contentType == "" {
		return mimetype.Detect(content).String()
	}

	return contentType
}

func isCompressible(contentType string) bool {
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	if strings.HasPrefix(mediaType, "text/") {
		return true
	}

	switch mediaType {
	case "application/javascript",
		"application/json",
		"application/manifest+json",
		"application/wasm",
		"application/xml",
		"image/svg+xml":
		return true
	default:
		return strings.HasSuffix(mediaType, "+json") ||
			strings.HasSuffix(mediaType, "+xml")
	}
}

func etag(content []byte) string {
	sum := sha256.Sum256(content)

	return `"` + hex.EncodeToString(sum[:]) + `"`
}

// Register mounts the server at mountPath. The path must start with a slash.
func (s *Server) Register(router fiber.Router, mountPath string) error {
	if mountPath == "" || !strings.HasPrefix(mountPath, "/") {
		return fmt.Errorf("static mount path must start with a slash")
	}

	mountPath = strings.TrimSuffix(mountPath, "/")
	if mountPath == "" {
		return fmt.Errorf("static mount path cannot be root")
	}

	handler := s.handler(mountPath)
	router.All(mountPath, handler)
	router.All(mountPath+"/*", handler)

	return nil
}

func (s *Server) handler(mountPath string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if ctx.Method() != http.MethodGet && ctx.Method() != http.MethodHead {
			return ctx.SendStatus(fiber.StatusMethodNotAllowed)
		}

		requestPath := ctx.Path()
		if requestPath == mountPath+"/" ||
			requestPath == mountPath+"/"+indexFile {
			return ctx.Redirect().
				Status(fiber.StatusPermanentRedirect).
				To(mountPath)
		}

		filename := strings.TrimPrefix(requestPath, mountPath+"/")
		if requestPath == mountPath {
			filename = indexFile
		}

		item, ok := s.files[filename]
		if !ok {
			if path.Ext(filename) != "" {
				return ctx.SendStatus(fiber.StatusNotFound)
			}

			item = s.files[indexFile]
		}

		return send(ctx, item)
	}
}

func send(ctx fiber.Ctx, item file) error {
	current, acceptable := selectEncoding(
		item,
		ctx.Get(fiber.HeaderAcceptEncoding),
	)
	if !acceptable {
		return ctx.SendStatus(fiber.StatusNotAcceptable)
	}

	rangeHeader := ctx.Get(fiber.HeaderRange)
	useRange := rangeHeader != "" && ifRangeMatches(
		ctx.Get(fiber.HeaderIfRange),
		item.identityETag,
	)

	if useRange {
		current = representation{
			content: item.identity,
			etag:    item.identityETag,
		}
	}

	setHeaders(ctx, item.contentType, current)

	if etagMatches(ctx.Get(fiber.HeaderIfNoneMatch), current.etag) {
		return ctx.SendStatus(fiber.StatusNotModified)
	}

	if useRange {
		ranges, err := ctx.Req().Range(int64(len(item.identity)))
		if err != nil {
			ctx.Res().Set(
				fiber.HeaderContentRange,
				"bytes */"+strconv.Itoa(len(item.identity)),
			)

			return ctx.SendStatus(fiber.StatusRequestedRangeNotSatisfiable)
		}

		return sendRanges(ctx, item, ranges.Ranges)
	}

	ctx.Res().Set(fiber.HeaderContentLength, strconv.Itoa(len(current.content)))

	if ctx.Method() == http.MethodHead {
		return nil
	}

	return ctx.Send(current.content)
}

func setHeaders(ctx fiber.Ctx, contentType string, item representation) {
	ctx.Res().Set(fiber.HeaderAcceptRanges, "bytes")
	ctx.Res().Set(fiber.HeaderCacheControl, "no-cache")
	ctx.Res().Set(fiber.HeaderContentType, contentType)
	ctx.Res().Set(fiber.HeaderETag, item.etag)
	ctx.Res().Set(fiber.HeaderVary, fiber.HeaderAcceptEncoding)
	ctx.Res().Set("X-Content-Type-Options", "nosniff")

	if item.encoding != "" {
		ctx.Res().Set(fiber.HeaderContentEncoding, item.encoding)
	}
}

func sendRanges(ctx fiber.Ctx, item file, ranges []fiber.RangeSet) error {
	if len(ranges) == 1 {
		return sendRange(ctx, item, ranges[0])
	}

	return sendMultiRange(ctx, item, ranges)
}

func sendRange(ctx fiber.Ctx, item file, byteRange fiber.RangeSet) error {
	content := item.identity[byteRange.Start : byteRange.End+1]
	ctx.Res().Set(
		fiber.HeaderContentRange,
		fmt.Sprintf(
			"bytes %d-%d/%d",
			byteRange.Start,
			byteRange.End,
			len(item.identity),
		),
	)
	ctx.Res().Set(fiber.HeaderContentLength, strconv.Itoa(len(content)))
	ctx.Status(fiber.StatusPartialContent)

	if ctx.Method() == http.MethodHead {
		return nil
	}

	return ctx.Send(content)
}

func sendMultiRange(ctx fiber.Ctx, item file, ranges []fiber.RangeSet) error {
	boundary := "platform-" + strings.Trim(item.identityETag, `"`)

	var body bytes.Buffer

	for _, byteRange := range ranges {
		fmt.Fprintf(&body, "--%s\r\n", boundary)
		fmt.Fprintf(&body, "Content-Type: %s\r\n", item.contentType)
		fmt.Fprintf(
			&body,
			"Content-Range: bytes %d-%d/%d\r\n\r\n",
			byteRange.Start,
			byteRange.End,
			len(item.identity),
		)
		body.Write(item.identity[byteRange.Start : byteRange.End+1])
		body.WriteString("\r\n")
	}

	fmt.Fprintf(&body, "--%s--\r\n", boundary)
	ctx.Res().Set(
		fiber.HeaderContentType,
		"multipart/byteranges; boundary="+boundary,
	)
	ctx.Res().Set(fiber.HeaderContentLength, strconv.Itoa(body.Len()))
	ctx.Status(fiber.StatusPartialContent)

	if ctx.Method() == http.MethodHead {
		return nil
	}

	return ctx.Send(body.Bytes())
}

func selectEncoding(item file, acceptEncoding string) (representation, bool) {
	qualities := encodingQualities(acceptEncoding)
	brotliQuality := quality(qualities, "br")
	gzipQuality := quality(qualities, "gzip")

	if item.compressible && brotliQuality > 0 && brotliQuality >= gzipQuality {
		return representation{
			content:  item.brotli,
			encoding: "br",
			etag:     item.brotliETag,
		}, true
	}

	if item.compressible && gzipQuality > 0 {
		return representation{
			content:  item.gzip,
			encoding: "gzip",
			etag:     item.gzipETag,
		}, true
	}

	if identityQuality(qualities) <= 0 {
		return representation{}, false
	}

	return representation{content: item.identity, etag: item.identityETag}, true
}

func encodingQualities(header string) map[string]float64 {
	qualities := make(map[string]float64)

	for value := range strings.SplitSeq(header, ",") {
		parts := strings.Split(value, ";")
		name := strings.TrimSpace(strings.ToLower(parts[0]))

		if name == "" {
			continue
		}

		quality := 1.0

		for _, parameter := range parts[1:] {
			parameter = strings.TrimSpace(parameter)
			if !strings.HasPrefix(parameter, "q=") {
				continue
			}

			if _, err := fmt.Sscanf(parameter, "q=%f", &quality); err != nil {
				quality = 0
			}
		}

		qualities[name] = quality
	}

	return qualities
}

func quality(qualities map[string]float64, encoding string) float64 {
	if quality, ok := qualities[encoding]; ok {
		return quality
	}

	return qualities["*"]
}

func identityQuality(qualities map[string]float64) float64 {
	if quality, ok := qualities["identity"]; ok {
		return quality
	}

	if quality, ok := qualities["*"]; ok {
		return quality
	}

	return 1
}

func etagMatches(header, etag string) bool {
	for value := range strings.SplitSeq(header, ",") {
		value = strings.TrimSpace(value)
		if value == "*" || value == etag {
			return true
		}
	}

	return false
}

func ifRangeMatches(header, etag string) bool {
	return header == "" || strings.TrimSpace(header) == etag
}
