package system

import (
	"context"

	etp "github.com/elum-utils/go-etp"
)

// DisconnectHandler runs once for a peer that completed the ETP handshake.
func DisconnectHandler(context.Context, *etp.Peer, error) {}
