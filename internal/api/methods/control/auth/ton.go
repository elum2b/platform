package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type TONRequest struct {
	Address         string                      `json:"address"           validate:"required"`
	Network         string                      `json:"network"`
	PublicKey       string                      `json:"public_key"`
	WalletStateInit string                      `json:"wallet_state_init"`
	Proof           controlauth.TONConnectProof `json:"proof"             validate:"required"`
	InviteToken     string                      `json:"invite_token"`
	BindToIP        bool                        `json:"bind_to_ip"`
}

var (
	tonKey         = "control.auth.ton"
	tonDescription = `
Authenticates an account through a TON Connect proof.`
)

// TON authenticates an account through TON Connect proof.
var TON = adapter.Method[TONRequest, AuthResponse]{
	Key:         tonKey,
	Description: tonDescription,
	Transports:  adapter.HTTP,
	Method:      fiber.MethodPost,
	Handler: func(
		ctx *adapter.Context,
		data TONRequest,
	) (AuthResponse, error) {
		metadata := identityMetadata(ctx, data.InviteToken, data.BindToIP)

		identity, err := controlauth.TONConnect(
			ctx.Context,
			controlauth.TONConnectAuthParams{
				Address:           data.Address,
				Network:           data.Network,
				PublicKey:         data.PublicKey,
				WalletStateInit:   data.WalletStateInit,
				Proof:             data.Proof,
				PayloadSecret:     config.ControlAuthTONPayloadSecret,
				ExpectedDomain:    config.ControlAuthTONDomain,
				ExpectedNetwork:   config.ControlAuthTONNetwork,
				AllowNativeDomain: config.ControlAuthTONAllowNativeDomain,
				InviteToken:       metadata.InviteToken,
				IP:                metadata.IP,
				UserAgent:         metadata.UserAgent,
				BindToIP:          metadata.BindToIP,
				ExpiresAt:         metadata.ExpiresAt,
				MaxAge:            config.ControlAuthTONMaxAge,
			},
		)
		if err != nil {
			return AuthResponse{}, err
		}

		return complete(ctx, identity)
	},
}
