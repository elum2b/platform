package methods

import (
	"github.com/elum2b/platform/internal/api/methods/calendar"
	"github.com/elum2b/platform/internal/api/methods/control"
	"github.com/elum2b/platform/internal/api/methods/cpa"
	"github.com/elum2b/platform/internal/api/methods/payment"
	"github.com/elum2b/platform/internal/api/methods/promo"
	"github.com/elum2b/platform/internal/api/methods/reference"
	"github.com/elum2b/platform/internal/api/methods/system"
	"github.com/elum2b/platform/internal/api/methods/tasks"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all platform API methods.
func Register(registry adapter.Registry) {
	system.Register(registry)
	control.Register(registry)
	cpa.Register(registry)
	promo.Register(registry)
	payment.Register(registry)
	calendar.Register(registry)
	reference.Register(registry)
	tasks.Register(registry)
}
