package catalog

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveLocalizationRequest struct {
	WorkspaceID     string `json:"workspace_id"      validate:"required,uuid"`
	Locale          string `json:"locale"            validate:"required,max=255"`
	LocalizationKey string `json:"localization_key"  validate:"required,max=255"`
	Value           string `json:"value"             validate:"required"`
}

var (
	saveLocalizationKey         = "payment.catalog.save_localization"
	saveLocalizationDescription = `
Saves a localization as part of catalog management. Requires the
'payment.catalog.save_localization' permission in the target workspace.`
)

var SaveLocalization = adapter.Method[SaveLocalizationRequest, struct{}]{
	Key:         saveLocalizationKey,
	Description: saveLocalizationDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(saveLocalizationKey)},
	Handler: func(ctx *adapter.Context, d SaveLocalizationRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.SaveLocalization(
			ctx.Context,
			padm.SaveLocalizationParams{
				WorkspaceID:     d.WorkspaceID,
				Locale:          d.Locale,
				LocalizationKey: d.LocalizationKey,
				Value:           d.Value,
			})
	},
}
