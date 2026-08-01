package socket

import (
	etp "github.com/elum-utils/go-etp"
	etpfiber "github.com/elum-utils/go-etp/adapters/fiber"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/api/socket_api/controllers/control"
	"github.com/elum2b/platform/internal/api/socket_api/system"
	"github.com/elum2b/platform/internal/config"
)

// Init registers the ETP WebSocket endpoint on the shared Fiber application.
// Event handlers are registered on the returned ETP app.
func Init(app *fiber.App) {
	adapter := etpfiber.Adapter{StrictFrameBoundary: true}

	sessionConfig := etp.DefaultServerConfig()

	sessionConfig.RateLimit.MaxFramesPerSecond = 200    // 200 rps
	sessionConfig.RateLimit.MaxBytesPerSecond = 5 << 20 // 5 MiB/s

	ws := etp.New(etp.Config{Session: sessionConfig})

	ws.OnAuth(system.AuthHandler)
	ws.OnConnect(system.ConnectHandler)
	ws.OnDisconnect(system.DisconnectHandler)
	ws.OnNotFound(system.NotFoundHandler)

	ws.OnError(system.ErrorHandler)
	ws.OnProtocolEvent(system.ProtocolEventHandler)
	ws.OnProgress(system.ProgressHandler)

	control.Register(ws)

	handler := adapter.Handler(ws)

	ws.Compile()

	app.Get("/ws", func(ctx fiber.Ctx) error {
		ctx.SetContext(
			system.WithSessionToken(
				ctx.Context(),
				ctx.Cookies(config.ControlAuthCookieName),
			),
		)

		return handler(ctx)
	})
}
