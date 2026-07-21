package system

import (
	"context"

	etp "github.com/elum-utils/go-etp"
)

// ProtocolEventHandler receives protocol violations and transport-level events.
func ProtocolEventHandler(context.Context, *etp.Peer, etp.ProtocolEvent) {}
