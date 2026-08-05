package reward

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string  `json:"workspace_id"   validate:"required,uuid"`
	CPAID       string  `json:"cpa_id"         validate:"required,max=255"`
	Key         string  `json:"key"            validate:"required,max=255"`
	Type        string  `json:"type"           validate:"required,max=255"`
	Quantity    int64   `json:"quantity"`
	Scale       uint16  `json:"scale"`
	Unit        *string `json:"unit,omitempty"`
}

var (
	upsertKey         = "cpa.reward.upsert"
	upsertDescription = `
Creates or updates a reward for a CPA offer. Requires the 'cpa.reward.upsert'
permission in the target workspace.`
)

// Upsert exposes the CPA reward upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.CPA.Admin.UpsertReward(
			ctx.Context,
			cpaadmin.UpsertRewardParams{
				WorkspaceID: data.WorkspaceID,
				CPAID:       data.CPAID,
				Key:         data.Key,
				Type:        data.Type,
				Quantity:    data.Quantity,
				Scale:       data.Scale,
				Unit:        data.Unit,
			},
		)

		return struct{}{}, err
	},
}
