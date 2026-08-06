package mcp

import (
	"context"
	"net/http"
	"strings"

	controlinternal "github.com/elum2b/services/control/service/internalapi"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	mcpserver "github.com/modelcontextprotocol/go-sdk/mcp"

	versionutils "github.com/elum2b/platform/internal/utils/version"
)

type principalContextKey struct{}

// RegisterFunc registers all MCP tools.
type RegisterFunc func(Router)

// Init registers the authenticated stateless Streamable HTTP MCP endpoint.
func Init(app fiber.Router, register RegisterFunc) {
	app.All(
		"/mcp",
		adaptor.HTTPHandler(HTTPHandler(register)),
	)
}

// HTTPHandler returns the authenticated MCP Streamable HTTP handler.
func HTTPHandler(register RegisterFunc) http.Handler {
	transport := mcpserver.NewStreamableHTTPHandler(
		func(*http.Request) *mcpserver.Server {
			server := server()
			router := NewRouter(server)
			group := router.Group()

			_ = group.Use("*", Authenticated)

			register(group)
			router.Compile()

			return server
		},
		&mcpserver.StreamableHTTPOptions{
			Stateless:      true,
			JSONResponse:   true,
			SessionTimeout: 0,
		},
	)

	return http.HandlerFunc(
		func(response http.ResponseWriter, request *http.Request) {
			transport.ServeHTTP(
				response,
				request.WithContext(
					withToken(request.Context(), bearerToken(request)),
				),
			)
		},
	)
}

// PrincipalFromContext returns the validated MCP token principal.
// Tool handlers use it to identify the account on whose behalf they run.
func PrincipalFromContext(
	ctx context.Context,
) (controlinternal.MCPPrincipal, bool) {
	principal, ok := ctx.Value(principalContextKey{}).(controlinternal.MCPPrincipal)

	return principal, ok
}

// WithPrincipal adds a validated MCP token principal to a context.
func WithPrincipal(
	ctx context.Context,
	principal controlinternal.MCPPrincipal,
) context.Context {
	return context.WithValue(ctx, principalContextKey{}, principal)
}

// server creates an MCP server with the platform implementation metadata.
func server() *mcpserver.Server {
	return mcpserver.NewServer(
		&mcpserver.Implementation{
			Name:    "elum2b-platform",
			Version: versionutils.Current(),
		},
		nil,
	)
}

func bearerToken(request *http.Request) string {
	parts := strings.Fields(request.Header.Get("Authorization"))
	if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
		return parts[1]
	}

	return request.URL.Query().Get("mcpToken")
}
