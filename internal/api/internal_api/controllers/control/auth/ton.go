package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	httputils "github.com/elum2b/platform/internal/utils/http"
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

type TONChallengeResponse struct {
	Payload string `json:"payload"`
}

func TONChallenge(path string, app fiber.Router) {
	app.Get(path, func(ctx fiber.Ctx) error {
		payload, err := controlauth.TONConnectPayload(
			config.ControlAuthTONPayloadSecret,
			config.ControlAuthTONMaxAge,
		)
		if err != nil {
			return httputils.Error(ctx, err)
		}

		return httputils.Respond(ctx, TONChallengeResponse{Payload: payload})
	})
}

func TON(path string, app fiber.Router) {
	app.Post(path, func(ctx fiber.Ctx) error {
		if !controlReady() {
			return httputils.Error(ctx, serviceerrors.ErrNotReady)
		}

		data := new(TONRequest)
		if !httputils.Decode(ctx, data) {
			return httputils.Error(ctx, serviceerrors.ErrInvalidFields)
		}

		metadata := identityMetadata(ctx, data.InviteToken, data.BindToIP)

		identity, err := controlauth.TONConnect(
			ctx.Context(),
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
			return httputils.Error(ctx, err)
		}

		return complete(ctx, identity)
	})
}
