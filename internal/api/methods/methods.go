package methods

import (
	"github.com/elum2b/platform/internal/api/methods/control"
	"github.com/elum2b/platform/internal/api/methods/cpa"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all platform API methods.
func Register(registry adapter.Registry) {
	control.Register(registry)
	cpa.Register(registry)
}
