package limitrequest

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CancelRequest struct {
	RequestID string `json:"request_id" validate:"required,uuid"`
}

type CancelResponse struct {
	Affected int64 `json:"affected"`
}

var (
	cancelKey         = "control.account.limit_request.cancel"
	cancelDescription = `
Cancels a pending limit request owned by the account.`
)

// Cancel cancels an account-owned pending limit request.
var Cancel = adapter.Method[CancelRequest, CancelResponse]{
	Key:         cancelKey,
	Description: cancelDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, data CancelRequest) (CancelResponse, error) {
		affected, err := services.Control.Admin.CancelLimitRequest(
			ctx.Context,
			ctx.AccountID,
			data.RequestID,
		)
		if err != nil {
			return CancelResponse{}, err
		}

		return CancelResponse{Affected: affected}, nil
	},
}
