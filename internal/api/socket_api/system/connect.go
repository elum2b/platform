package system

import (
	"context"

	etp "github.com/elum-utils/go-etp"
)

// ConnectHandler runs after successful ETP authentication and handshake.
func ConnectHandler(context.Context, *etp.Peer) error {
	return nil
}
