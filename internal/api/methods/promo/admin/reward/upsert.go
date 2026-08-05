package reward

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string  `json:"workspace_id"   validate:"required,uuid"`
	PromoID     uint64  `json:"promo_id"       validate:"required,min=1"`
	Key         string  `json:"key"            validate:"required,max=255"`
	Type        string  `json:"type"           validate:"required,max=255"`
	Quantity    int64   `json:"quantity"`
	Scale       uint16  `json:"scale"`
	Unit        *string `json:"unit,omitempty"`
}

var (
	upsertKey         = "promo.reward.upsert"
	upsertDescription = `
Creates or updates a reward for a promo. Requires the 'promo.reward.upsert'
permission in the target workspace.`
)

// Upsert exposes the promo reward upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Promo.Admin.UpsertReward(
			ctx.Context,
			promoadmin.SaveRewardParams{
				WorkspaceID: data.WorkspaceID,
				PromoID:     data.PromoID,
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
