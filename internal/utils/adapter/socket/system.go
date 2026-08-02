package socket

import (
	"context"

	etp "github.com/elum-utils/go-etp"
)

func connectHandler(context.Context, *etp.Peer) error { return nil }

func disconnectHandler(context.Context, *etp.Peer, error) {}

func notFoundHandler(*etp.Context) error { return etp.ErrRouteNotFound }

func errorHandler(*etp.Context, error) {}

func protocolEventHandler(context.Context, *etp.Peer, etp.ProtocolEvent) {}

func progressHandler(context.Context, *etp.Peer, etp.Progress) {}
