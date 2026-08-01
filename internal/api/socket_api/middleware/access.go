package middleware

import (
	"strings"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/internalapi"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type workspaceAccessRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

func GlobalAccess(method string) etp.Middleware {
	return func(next etp.Handler) etp.Handler {
		return func(ctx *etp.Context) error {
			accountID, err := accessAccountID(ctx)
			if err != nil {
				return err
			}

			if services.Control == nil || !services.Control.IsReady() {
				return serviceerrors.ErrNotReady
			}

			allowed, err := services.Control.Internal.CheckGlobalAccess(
				ctx,
				internalapi.GlobalAccessRequest{
					AccountID: accountID,
					MethodKey: method,
				},
			)
			if err != nil {
				return err
			}

			if !allowed {
				return serviceerrors.ErrForbidden
			}

			return next(ctx)
		}
	}
}

func WorkspaceAccess(method string) etp.Middleware {
	return func(next etp.Handler) etp.Handler {
		return func(ctx *etp.Context) error {
			accountID, err := accessAccountID(ctx)
			if err != nil {
				return err
			}

			if services.Control == nil || !services.Control.IsReady() {
				return serviceerrors.ErrNotReady
			}

			data := new(workspaceAccessRequest)
			if !socketutils.Decode(ctx, data) {
				return serviceerrors.ErrInvalidFields
			}

			allowed, err := services.Control.Internal.CheckWorkspaceAccess(
				ctx,
				internalapi.WorkspaceAccessRequest{
					AccountID:   accountID,
					WorkspaceID: data.WorkspaceID,
					MethodKey:   method,
				},
			)
			if err != nil {
				return err
			}

			if !allowed {
				return serviceerrors.ErrForbidden
			}

			return next(ctx)
		}
	}
}

func accessAccountID(ctx *etp.Context) (string, error) {
	if ctx == nil || ctx.Peer == nil {
		return "", serviceerrors.ErrUnauthorized
	}

	accountID := strings.TrimSpace(ctx.Peer.Identity().UserID)
	if accountID == "" {
		return "", serviceerrors.ErrUnauthorized
	}

	return accountID, nil
}
