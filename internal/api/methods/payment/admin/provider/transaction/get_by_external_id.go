package transaction

import (
	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetByExternalIDRequest struct {
	WorkspaceID           string `json:"workspace_id"            validate:"required,uuid"`
	ProviderCode          string `json:"provider_code"           validate:"required,max=255"`
	Network               string `json:"network"                 validate:"required,max=255"`
	SourceKey             string `json:"source_key"              validate:"required,max=255"`
	ExternalTransactionID string `json:"external_transaction_id" validate:"required,max=255"`
}

type GetByExternalIDResponse struct {
	Transaction padm.ProviderTransactionModel `json:"transaction"`
}

var (
	getByExternalIDKey         = "payment.provider_transaction.get_by_external_id"
	getByExternalIDDescription = `
Returns a provider transaction by external ID. Requires the
'payment.provider_transaction.get_by_external_id' permission in the target workspace.`
)

var GetByExternalID = adapter.Method[GetByExternalIDRequest, GetByExternalIDResponse]{
	Key:         getByExternalIDKey,
	Description: getByExternalIDDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getByExternalIDKey),
	},
	Handler: func(ctx *adapter.Context, d GetByExternalIDRequest) (GetByExternalIDResponse, error) {
		v, err := services.Payment.Admin.GetProviderTransactionByExternalID(
			ctx.Context,
			d.WorkspaceID,
			d.ProviderCode,
			d.Network,
			d.SourceKey,
			d.ExternalTransactionID,
		)

		return GetByExternalIDResponse{Transaction: v}, err
	},
}
