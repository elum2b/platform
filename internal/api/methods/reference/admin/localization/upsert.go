package localization

import (
	refadmin "github.com/elum2b/services/reference/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ItemKey     string `json:"item_key"     validate:"required,max=255"`
	Locale      string `json:"locale"       validate:"required,max=32"`
	Title       string `json:"title"        validate:"required"`
	Description string `json:"description"  validate:"required"`
}

var (
	upsertKey         = "reference.localization.upsert"
	upsertDescription = `
Creates or updates a reference item localization. Requires the
'reference.localization.upsert' permission in the target workspace.`
)

// Upsert exposes the reference localization upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.Reference.Admin.UpsertLocalization(
			ctx.Context,
			refadmin.SaveLocalizationParams{
				WorkspaceID: data.WorkspaceID,
				ItemKey:     data.ItemKey,
				Locale:      data.Locale,
				Title:       data.Title,
				Description: data.Description,
			},
		)

		return struct{}{}, err
	},
}
