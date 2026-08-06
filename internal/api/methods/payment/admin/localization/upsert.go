package localization

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID     string `json:"workspace_id"     validate:"required,uuid"`
	Locale          string `json:"locale"           validate:"required,max=255"`
	LocalizationKey string `json:"localization_key" validate:"required,max=255"`
	Value           string `json:"value"            validate:"required"`
}

var (
	upsertKey         = "payment.localization.upsert"
	upsertDescription = `
Creates or updates a localization. Requires the 'payment.localization.upsert'
permission in the target workspace.`
)

var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, d UpsertRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.UpsertLocalization(
			ctx.Context,
			padm.LocalizationUpsertParams{
				WorkspaceID:     d.WorkspaceID,
				Locale:          d.Locale,
				LocalizationKey: d.LocalizationKey,
				Value:           d.Value,
			})
	},
}
