package wallet

import (
	padm "github.com/elum2b/services/payment/service/admin"
	"github.com/elum2b/services/payment/tonconnect"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveRequest struct {
	WorkspaceID      string  `json:"workspace_id"                 validate:"required,uuid"`
	Network          string  `json:"network"                      validate:"required,max=255"`
	WalletAddress    string  `json:"wallet_address"               validate:"required"`
	NetworkConfigURL *string `json:"network_config_url,omitempty"`
	ManifestURL      string  `json:"manifest_url"                 validate:"required"`
	ManifestName     string  `json:"manifest_name"                validate:"required"`
	ManifestIconURL  string  `json:"manifest_icon_url"            validate:"required"`
	TermsOfUseURL    *string `json:"terms_of_use_url,omitempty"`
	PrivacyPolicyURL *string `json:"privacy_policy_url,omitempty"`
	IsEnabled        bool    `json:"is_enabled"`
}

var (
	saveKey         = "payment.ton_wallet.save"
	saveDescription = `
Saves a TON wallet configuration. Requires the 'payment.ton_wallet.save'
permission in the target workspace.`
)

var Save = adapter.Method[SaveRequest, struct{}]{
	Key:         saveKey,
	Description: saveDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(saveKey)},
	Handler: func(ctx *adapter.Context, d SaveRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.SaveTONWallet(
			ctx.Context,
			padm.TONWalletUpsertParams{
				WorkspaceID:      d.WorkspaceID,
				Network:          d.Network,
				WalletAddress:    d.WalletAddress,
				NetworkConfigURL: d.NetworkConfigURL,
				Manifest: tonconnect.Manifest{
					URL:              d.ManifestURL,
					Name:             d.ManifestName,
					IconURL:          d.ManifestIconURL,
					TermsOfUseURL:    d.TermsOfUseURL,
					PrivacyPolicyURL: d.PrivacyPolicyURL,
				},
				IsEnabled: d.IsEnabled,
			})
	},
}
