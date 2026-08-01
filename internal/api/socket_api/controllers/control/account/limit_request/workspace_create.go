package limitrequest

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	controlmodel "github.com/elum2b/services/control/model"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type WorkspaceLimitRequest struct {
	RequestedLimit int32  `json:"requested_limit" validate:"required,min=1"`
	Reason         string `json:"reason"          validate:"required,max=1000"`
}

type WorkspaceLimitResponse struct {
	ID             string                          `json:"id"`
	Kind           string                          `json:"kind"`
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

func WorkspaceCreate(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		data := new(WorkspaceLimitRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		request, err := services.Control.Admin.RequestWorkspaceLimit(
			ctx,
			ctx.Peer.Identity().UserID,
			data.RequestedLimit,
			data.Reason,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, WorkspaceLimitResponse{
			ID:             request.ID,
			Kind:           string(request.Kind),
			AccountID:      request.AccountID,
			WorkspaceID:    request.WorkspaceID,
			CurrentLimit:   request.CurrentLimit,
			RequestedLimit: request.RequestedLimit,
			ApprovedLimit:  request.ApprovedLimit,
			Reason:         request.Reason,
			Status:         request.Status,
			RequestedBy:    request.RequestedBy,
			ReviewedBy:     request.ReviewedBy,
			ReviewComment:  request.ReviewComment,
			CreatedAt:      request.CreatedAt,
			ReviewedAt:     request.ReviewedAt,
		})
	})
}
