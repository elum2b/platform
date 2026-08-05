package event

import (
	cpamodel "github.com/elum2b/services/cpa/model"
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string                       `json:"workspace_id"         validate:"required,uuid"`
	CPAID       string                       `json:"cpa_id"               validate:"required,max=255"`
	EventType   cpamodel.AssignmentEventType `json:"event_type,omitempty" validate:"omitempty,oneof=issued completed"`
	Limit       int32                        `json:"limit,omitempty"      validate:"omitempty,min=1,max=100"`
	Offset      int32                        `json:"offset,omitempty"     validate:"min=0"`
}
type ListResponse struct {
	Events []cpaadmin.AssignmentEventModel `json:"events"`
}

var (
	listKey         = "cpa.assignment.event.list"
	listDescription = `
Lists lifecycle events of CPA assignments. Requires the
'cpa.assignment.event.list' permission in the target workspace.`
)

// List exposes the CPA assignment event listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.CPA.Admin.ListAssignmentEvents(
			ctx.Context,
			cpaadmin.AssignmentEventListParams{
				WorkspaceID: data.WorkspaceID,
				CPAID:       data.CPAID,
				EventType:   data.EventType,
				Page: cpaadmin.Page{
					Limit:  data.Limit,
					Offset: data.Offset,
				},
			},
		)

		return ListResponse{Events: values}, err
	},
}
