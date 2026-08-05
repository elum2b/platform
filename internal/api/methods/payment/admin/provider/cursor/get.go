package cursor

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID  string `json:"workspace_id"  validate:"required,uuid"`
	ProviderCode string `json:"provider_code" validate:"required,max=255"`
	Network      string `json:"network"       validate:"required,max=255"`
	SourceKey    string `json:"source_key"    validate:"required,max=255"`
}

type GetResponse struct {
	Cursor padm.ProviderCursorModel `json:"cursor"`
}

var (
	getKey         = "payment.provider_cursor.get"
	getDescription = `
Returns a provider cursor. Requires the 'payment.provider_cursor.get'
permission in the target workspace.`
)

var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(getKey)},
	Handler: func(ctx *adapter.Context, d GetRequest) (GetResponse, error) {
		v, err := services.Payment.Admin.GetProviderCursor(
			ctx.Context, d.WorkspaceID, d.ProviderCode, d.Network, d.SourceKey)
		return GetResponse{Cursor: v}, err
	},
}
