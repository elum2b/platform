package cursor

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID  string `json:"workspace_id"  validate:"required,uuid"`
	ProviderCode string `json:"provider_code" validate:"required,max=255"`
	Network      string `json:"network"       validate:"required,max=255"`
	SourceKey    string `json:"source_key"    validate:"required,max=255"`
	CursorValue  string `json:"cursor_value"  validate:"required,max=255"`
}

type UpsertResponse struct {
	Affected int64 `json:"affected"`
}

var (
	upsertKey         = "payment.provider_cursor.upsert"
	upsertDescription = `
Creates or updates a provider cursor. Requires the
'payment.provider_cursor.upsert' permission in the target workspace.`
)

var Upsert = adapter.Method[UpsertRequest, UpsertResponse]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, d UpsertRequest) (UpsertResponse, error) {
		a, err := services.Payment.Admin.UpsertProviderCursor(
			ctx.Context,
			padm.ProviderCursorUpsertParams{
				WorkspaceID:  d.WorkspaceID,
				ProviderCode: d.ProviderCode,
				Network:      d.Network,
				SourceKey:    d.SourceKey,
				CursorValue:  d.CursorValue,
			})
		return UpsertResponse{Affected: a}, err
	},
}
