package system

import (
	"context"

	etp "github.com/elum-utils/go-etp"
)

// AuthHandler validates the ETP authentication payload before a peer is
// considered connected. It rejects every request until application auth is set.
func AuthHandler(context.Context, etp.AuthRequest) (etp.AuthResult, error) {
	return etp.AuthResult{
		OK:     false,
		Reason: "authentication is not configured",
	}, nil
}
