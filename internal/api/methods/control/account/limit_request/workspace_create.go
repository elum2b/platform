package limitrequest

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type WorkspaceCreateRequest struct {
	RequestedLimit int32  `json:"requested_limit" validate:"required,min=1"`
	Reason         string `json:"reason"          validate:"required,max=1000"`
}

type WorkspaceCreateResponse struct {
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

var (
	workspaceCreateKey         = "control.account.workspace_limit.request"
	workspaceCreateDescription = `
Requests a larger workspace ownership limit for the account.`
)

// WorkspaceCreate requests a larger workspace limit for the account.
var WorkspaceCreate = adapter.Method[WorkspaceCreateRequest, WorkspaceCreateResponse]{
	Key:         workspaceCreateKey,
	Description: workspaceCreateDescription,
	Transports:  adapter.WS,
	Handler: func(
		ctx *adapter.Context,
		data WorkspaceCreateRequest,
	) (WorkspaceCreateResponse, error) {
		request, err := services.Control.Admin.RequestWorkspaceLimit(
			ctx.Context,
			ctx.AccountID,
			data.RequestedLimit,
			data.Reason,
		)
		if err != nil {
			return WorkspaceCreateResponse{}, err
		}

		return WorkspaceCreateResponse{
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
		}, nil
	},
}
