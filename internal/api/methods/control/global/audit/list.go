package audit

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"
	json "github.com/goccy/go-json"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	Limit    int32     `json:"limit,omitempty"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at,omitempty"`
	CursorID string    `json:"cursor_id,omitempty"`
}

type Item struct {
	ID          string                   `json:"id"`
	Scope       controladmin.AccessScope `json:"scope"`
	WorkspaceID string                   `json:"workspace_id"`
	ActorID     string                   `json:"actor_id"`
	MethodKey   string                   `json:"method_key"`
	TargetType  string                   `json:"target_type"`
	TargetID    string                   `json:"target_id"`
	Result      controlmodel.AuditResult `json:"result"`
	RequestID   string                   `json:"request_id"`
	BeforeData  json.RawMessage          `json:"before_data"`
	AfterData   json.RawMessage          `json:"after_data"`
	OccurredAt  time.Time                `json:"occurred_at"`
}

type ListResponse struct {
	Events []Item `json:"events"`
}

var (
	listKey         = "control.global.audit.list"
	listDescription = `
Lists global platform audit entries. Requires the
'control.global.audit.list' global permission.`
)

// List returns global audit entries.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess("control.global.audit.list"),
	},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		items, err := services.Control.Admin.ListGlobalAudit(
			ctx.Context,
			controladmin.Page{
				Limit: data.Limit, CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		events := make([]Item, 0, len(items))
		for _, item := range items {
			events = append(events, Item{
				ID: item.ID, Scope: item.Scope, WorkspaceID: item.WorkspaceID,
				ActorID: item.ActorID, MethodKey: item.MethodKey,
				TargetType: item.TargetType, TargetID: item.TargetID,
				Result: item.Result, RequestID: item.RequestID,
				BeforeData: item.BeforeData, AfterData: item.AfterData,
				OccurredAt: item.OccurredAt,
			})
		}

		return ListResponse{Events: events}, nil
	},
}
