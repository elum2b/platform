package middleware

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
)

func ControlReady(next etp.Handler) etp.Handler {
	return func(ctx *etp.Context) error {
		if services.Control == nil || !services.Control.IsReady() {
			return serviceerrors.ErrNotReady
		}

		return next(ctx)
	}
}
