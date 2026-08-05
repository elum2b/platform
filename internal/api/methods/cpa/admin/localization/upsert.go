package localization

import (
	cpaadmin "github.com/elum2b/services/cpa/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	CPAID       string `json:"cpa_id"       validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"  validate:"required"`
}

var (
	upsertKey         = "cpa.localization.upsert"
	upsertDescription = `
Creates or updates a CPA offer localization. Requires the
'cpa.localization.upsert' permission in the target workspace.`
)

// Upsert exposes the CPA localization upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.CPA.Admin.UpsertLocalization(
			ctx.Context,
			cpaadmin.UpsertLocalizationParams{
				WorkspaceID: data.WorkspaceID,
				CPAID:       data.CPAID,
				Locale:      data.Locale,
				Title:       data.Title,
				Description: data.Description,
			},
		)

		return struct{}{}, err
	},
}
