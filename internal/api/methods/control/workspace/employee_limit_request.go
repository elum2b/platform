package workspace

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type EmployeeLimitRequest struct {
	WorkspaceID    string `json:"workspace_id"     validate:"required,uuid"`
	RequestedLimit int32  `json:"requested_limit"  validate:"required,min=1"`
	Reason         string `json:"reason,omitempty" validate:"max=1000"`
}

type EmployeeLimitResponse struct {
	ID             string                          `json:"id"`
	Kind           controladmin.LimitKind          `json:"kind"`
	WorkspaceID    string                          `json:"workspace_id"`
	CurrentLimit   int32                           `json:"current_limit"`
	RequestedLimit int32                           `json:"requested_limit"`
	Reason         string                          `json:"reason"`
	Status         controlmodel.LimitRequestStatus `json:"status"`
	CreatedAt      time.Time                       `json:"created_at"`
}

var (
	employeeLimitKey         = "control.workspace.employee_limit.request"
	employeeLimitDescription = `
Creates a request to increase the employee limit of a workspace. Requires the
'control.workspace.employee_limit.request' permission in the target workspace.`
)

// EmployeeLimit requests a larger employee limit for a workspace.
var EmployeeLimit = adapter.Method[EmployeeLimitRequest, EmployeeLimitResponse]{
	Key:         employeeLimitKey,
	Description: employeeLimitDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess("control.workspace.employee_limit.request"),
	},
	Handler: func(
		ctx *adapter.Context,
		data EmployeeLimitRequest,
	) (EmployeeLimitResponse, error) {
		item, err := services.Control.Admin.RequestEmployeeLimit(
			ctx.Context,
			ctx.AccountID,
			data.WorkspaceID,
			data.RequestedLimit,
			data.Reason,
		)
		if err != nil {
			return EmployeeLimitResponse{}, err
		}

		return EmployeeLimitResponse{
			ID: item.ID, Kind: item.Kind, WorkspaceID: item.WorkspaceID,
			CurrentLimit:   item.CurrentLimit,
			RequestedLimit: item.RequestedLimit, Reason: item.Reason,
			Status: item.Status, CreatedAt: item.CreatedAt,
		}, nil
	},
}
