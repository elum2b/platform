package assignment

import (
	cpamodel "github.com/elum2b/services/cpa/model"
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID string                    `json:"workspace_id"     validate:"required,uuid"`
	CPAID       string                    `json:"cpa_id"           validate:"required,max=255"`
	Status      cpamodel.AssignmentStatus `json:"status,omitempty" validate:"omitempty,oneof=issued completed"`
	Limit       int32                     `json:"limit,omitempty"  validate:"omitempty,min=1,max=100"`
	Offset      int32                     `json:"offset,omitempty" validate:"min=0"`
}
type ListResponse struct {
	Assignments []cpaadmin.AssignmentModel `json:"assignments"`
}

var (
	listKey         = "cpa.assignment.list"
	listDescription = `
Lists assignments of a CPA offer with optional status filtering. Requires the
'cpa.assignment.list' permission in the target workspace.`
)

// List exposes the CPA assignment listing method.
var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.CPA.Admin.ListAssignments(
			ctx.Context,
			cpaadmin.AssignmentListParams{
				WorkspaceID: data.WorkspaceID,
				CPAID:       data.CPAID,
				Status:      data.Status,
				Page: cpaadmin.Page{
					Limit:  data.Limit,
					Offset: data.Offset,
				},
			},
		)

		return ListResponse{Assignments: values}, err
	},
}
