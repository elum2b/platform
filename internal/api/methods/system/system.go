package system

import adapter "github.com/elum2b/platform/internal/utils/adapter"

// Register registers all public system methods.
func Register(registry adapter.Registry) {
	Version.Register(registry)
}
