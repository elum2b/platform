package api

import (
	"context"
	"fmt"

	socket "github.com/elum2b/platform/internal/api/socket_api"
	"github.com/elum2b/platform/internal/config"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"

	"github.com/goccy/go-json"
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

		// Register WebSocket routes and handlers.
		socket.Init(app)

		// // Register public/external API routes.
		// external.Init(app)

		// // Register private/internal API routes.
		// internal.Init(app)

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
