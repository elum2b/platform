package socket

import (
	etp "github.com/elum-utils/go-etp"
	etpfiber "github.com/elum-utils/go-etp/adapters/fiber"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
)

// RegisterFunc registers platform methods on the authenticated ETP router.
type RegisterFunc func(etp.Router)

// Init registers the authenticated ETP WebSocket endpoint.
func Init(app *fiber.App, register RegisterFunc) {
	fiberAdapter := etpfiber.Adapter{StrictFrameBoundary: true}

	sessionConfig := etp.DefaultServerConfig()

	sessionConfig.RateLimit.MaxFramesPerSecond = 200
	sessionConfig.RateLimit.MaxBytesPerSecond = 5 << 20

	socket := etp.New(etp.Config{Session: sessionConfig})
	socket.OnAuth(authHandler)
	socket.OnConnect(connectHandler)
	socket.OnDisconnect(disconnectHandler)
	socket.OnNotFound(notFoundHandler)
	socket.OnError(errorHandler)
	socket.OnProtocolEvent(protocolEventHandler)
	socket.OnProgress(progressHandler)

	group := socket.Group()
	group.Use("*", Authenticated)
	register(group)

	handler := fiberAdapter.Handler(socket)
	socket.Compile()

	app.Get("/ws", func(ctx fiber.Ctx) error {
		ctx.SetContext(withSessionToken(
			ctx.Context(),
			ctx.Cookies(config.ControlAuthCookieName),
		))

		return handler(ctx)
	})
}
