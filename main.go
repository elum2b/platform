package main

import (
	"context"
	"time"

	"github.com/elum-utils/supervisor"

	"github.com/elum2b/platform/internal/services/calendar"
	"github.com/elum2b/platform/internal/services/control"
	"github.com/elum2b/platform/internal/services/cpa"
	"github.com/elum2b/platform/internal/services/payment"
	"github.com/elum2b/platform/internal/services/promo"
	"github.com/elum2b/platform/internal/services/reference"
	"github.com/elum2b/platform/internal/services/tasks"
)

func main() {

	// Create the root service supervisor.
	// The supervisor manages lifecycle, restarts,
	// graceful shutdown and fault isolation for services.
	s := supervisor.New(
		context.Background(),
		supervisor.Options{

			// Maximum number of restart attempts allowed
			// within RestartInterval before the service
			// is considered permanently failed.
			RestartLimit: 5,

			// Time window used to count restart attempts.
			// If RestartLimit is exceeded during this interval,
			// the supervisor stops restarting the service.
			RestartInterval: time.Minute,

			// Delay before attempting to restart
			// a crashed or failed service.
			RestartDelay: time.Second * 5,

			// Maximum time allowed for graceful shutdown
			// before force termination.
			ShutdownTimeout: time.Minute,
		},
	)

	// Register the control service.
	// It manages accounts, workspaces, access control,
	// identities and sessions.
	s.Add("control", control.Service())

	// Register the payment service.
	// It manages products, checkout, payment providers,
	// refunds and callbacks.
	s.Add("payment", payment.Service())

	// Register the calendar service.
	// It manages calendars, user progress, rewards
	// and completion callbacks.
	s.Add("calendar", calendar.Service())

	// Register the CPA service.
	// It manages offers, personal or shared codes,
	// rewards and activation callbacks.
	s.Add("cpa", cpa.Service())

	// Register the promo service.
	// It manages promo codes, redemption limits,
	// rewards and redemption callbacks.
	s.Add("promo", promo.Service())

	// Register the reference service.
	// It provides canonical item metadata,
	// localization and custom payloads.
	s.Add("reference", reference.Service())

	// Register the tasks service.
	// It manages task chains, progress, reward
	// claims and completion callbacks.
	s.Add("tasks", tasks.Service())

	// Start supervisor event loop and block
	// until shutdown or fatal failure occurs.
	s.Run()

}
