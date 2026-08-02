package api

import (
	"context"
	"fmt"

	etp "github.com/elum-utils/go-etp"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/elum2b/platform/internal/api/methods"
	"github.com/elum2b/platform/internal/config"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	httputils "github.com/elum2b/platform/internal/utils/adapter/http"
	mcputils "github.com/elum2b/platform/internal/utils/adapter/mcp"
	socketutils "github.com/elum2b/platform/internal/utils/adapter/socket"
)

func Service() func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// Create Fiber application instance.
		// Custom JSON encoder/decoder are used to improve
		// serialization performance and keep JSON behavior consistent.
		app := fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		})

		// Recover from panics inside HTTP handlers
		// and prevent the whole service from crashing.
		app.Use(recover.New())

		// Register WebSocket methods.
		socketutils.Init(app, func(router etp.Router) {
			methods.Register(adapter.Registry{Socket: router})
		})

		// Register MCP methods.
		mcputils.Init(app, func(router mcputils.Router) {
			methods.Register(adapter.Registry{MCP: router})
		})

		// Register private/internal HTTP methods.
		httputils.Init(app, func(router fiber.Router) {
			methods.Register(adapter.Registry{HTTP: router})
		})

		// Start HTTP server and bind it to configured host and port.
		// GracefulContext allows the supervisor to stop the server
		// cleanly during shutdown.
		return app.Listen(
			fmt.Sprintf(
				"%v:%v",
				config.Host,
				config.Port,
			),
			fiber.ListenConfig{
				GracefulContext:       ctx,
				ListenerNetwork:       "tcp4",
				DisableStartupMessage: false,
			})
	}
}
