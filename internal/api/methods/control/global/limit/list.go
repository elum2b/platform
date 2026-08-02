package limit

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	Status   controlmodel.LimitRequestStatus `json:"status,omitempty"`
	Limit    int32                           `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time                       `json:"cursor_at,omitempty"`
	CursorID string                          `json:"cursor_id,omitempty"`
}

type Item struct {
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

type ListResponse struct {
	Requests []Item `json:"requests"`
}

var (
	listKey         = "control.global.limit.list"
	listDescription = `
Lists account and workspace limit requests. Requires the
'control.global.limit.list' global permission.`
)

// List returns platform limit requests.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.limit.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListLimitRequests(
			ctx.Context,
			ctx.AccountID,
			data.Status,
			controladmin.Page{
				Limit: data.Limit, CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		requests := make([]Item, 0, len(items))
		for _, item := range items {
			requests = append(requests, mapItem(item))
		}

		return ListResponse{Requests: requests}, nil
	},
}

func mapItem(item controladmin.LimitRequestModel) Item {
	return Item{
		ID: item.ID, Kind: item.Kind, AccountID: item.AccountID,
		WorkspaceID: item.WorkspaceID, CurrentLimit: item.CurrentLimit,
		RequestedLimit: item.RequestedLimit, ApprovedLimit: item.ApprovedLimit,
		Reason: item.Reason, Status: item.Status, RequestedBy: item.RequestedBy,
		ReviewedBy: item.ReviewedBy, ReviewComment: item.ReviewComment,
		CreatedAt: item.CreatedAt, ReviewedAt: item.ReviewedAt,
	}
}
