package system

import (
	"context"

	etp "github.com/elum-utils/go-etp"
)

// ProgressHandler receives progress for incoming and outgoing transfer bodies.
func ProgressHandler(context.Context, *etp.Peer, etp.Progress) {}
