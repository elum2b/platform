package socket

import (
	etp "github.com/elum-utils/go-etp"
	etpfiber "github.com/elum-utils/go-etp/adapters/fiber"
	"github.com/elum2b/platform/internal/api/socket_api/system"
	"github.com/gofiber/fiber/v3"
)

// Init registers the ETP WebSocket endpoint on the shared Fiber application.
// Event handlers are registered on the returned ETP app.
func Init(app *fiber.App) {

	adapter := etpfiber.Adapter{StrictFrameBoundary: true}

	config := etp.DefaultServerConfig()

	config.RateLimit.MaxFramesPerSecond = 200    // 200 rps
	config.RateLimit.MaxBytesPerSecond = 5 << 20 // 5 MiB/s
	ws := etp.New(etp.Config{Session: config})

	ws.OnAuth(system.AuthHandler)
	ws.OnConnect(system.ConnectHandler)
	ws.OnDisconnect(system.DisconnectHandler)
	ws.OnNotFound(system.NotFoundHandler)

	ws.OnError(system.ErrorHandler)
	ws.OnProtocolEvent(system.ProtocolEventHandler)
	ws.OnProgress(system.ProgressHandler)

	app.Get("/ws", adapter.Handler(ws))

	
	


	ws.Compile()

}
