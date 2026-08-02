package limit

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ResolveRequest struct {
	RequestID     string `json:"request_id"        validate:"required,uuid"`
	Approved      bool   `json:"approved"`
	ApprovedLimit int32  `json:"approved_limit"    validate:"min=0"`
	Comment       string `json:"comment,omitempty" validate:"max=1000"`
}

type ResolveResponse struct {
	ID             string                          `json:"id"`
	Kind           controladmin.LimitKind          `json:"kind"`
	AccountID      string                          `json:"account_id"`
	WorkspaceID    string                          `json:"workspace_id"`
	CurrentLimit   int32                           `json:"current_limit"`
	RequestedLimit int32                           `json:"requested_limit"`
	ApprovedLimit  *int32                          `json:"approved_limit"`
	Reason         string                          `json:"reason"`
	Status         controlmodel.LimitRequestStatus `json:"status"`
	RequestedBy    string                          `json:"requested_by"`
	ReviewedBy     string                          `json:"reviewed_by"`
	ReviewComment  string                          `json:"review_comment"`
	CreatedAt      time.Time                       `json:"created_at"`
	ReviewedAt     *time.Time                      `json:"reviewed_at"`
}

var (
	resolveKey         = "control.global.limit.resolve"
	resolveDescription = `
Approves or rejects a platform limit request. Requires the
'control.global.limit.resolve' global permission.`
)

// Resolve approves or rejects a limit request.
var Resolve = adapter.Method[ResolveRequest, ResolveResponse]{
	Key:         resolveKey,
	Description: resolveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.limit.resolve"),
	},
	Handler: func(ctx *adapter.Context, data ResolveRequest) (ResolveResponse, error) {
		item, err := services.Control.Admin.ResolveLimitRequest(
			ctx.Context,
			controladmin.ResolveLimitRequestParams{
				ActorID: ctx.AccountID, RequestID: data.RequestID,
				Approved: data.Approved, ApprovedLimit: data.ApprovedLimit,
				Comment: data.Comment,
			},
		)
		if err != nil {
			return ResolveResponse{}, err
		}

		return ResolveResponse{
			ID: item.ID, Kind: item.Kind, AccountID: item.AccountID,
			WorkspaceID: item.WorkspaceID, CurrentLimit: item.CurrentLimit,
			RequestedLimit: item.RequestedLimit,
			ApprovedLimit:  item.ApprovedLimit, Reason: item.Reason,
			Status: item.Status, RequestedBy: item.RequestedBy,
			ReviewedBy: item.ReviewedBy, ReviewComment: item.ReviewComment,
			CreatedAt: item.CreatedAt, ReviewedAt: item.ReviewedAt,
		}, nil
	},
}
