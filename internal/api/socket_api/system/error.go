package system

import etp "github.com/elum-utils/go-etp"

// ErrorHandler receives errors returned by socket middleware and event handlers.
func ErrorHandler(*etp.Context, error) {}
