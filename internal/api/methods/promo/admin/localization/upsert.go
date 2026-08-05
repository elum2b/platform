package localization

import (
	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	PromoID     uint64 `json:"promo_id"     validate:"required,min=1"`
	Locale      string `json:"locale"       validate:"required,max=32"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"  validate:"required"`
}

var (
	upsertKey         = "promo.localization.upsert"
	upsertDescription = `
Creates or updates a promo localization. Requires the
'promo.localization.upsert' permission in the target workspace.`
)

// Upsert exposes the promo localization upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Promo.Admin.UpsertLocalization(
			ctx.Context,
			promoadmin.SaveLocalizationParams{
				WorkspaceID: data.WorkspaceID,
				PromoID:     data.PromoID,
				Locale:      data.Locale,
				Title:       data.Title,
				Description: data.Description,
			},
		)

		return struct{}{}, err
	},
}
