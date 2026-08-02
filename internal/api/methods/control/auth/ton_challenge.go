package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type TONChallengeResponse struct {
	Payload string `json:"payload"`
}

var (
	tonChallengeKey         = "control.auth.ton.challenge"
	tonChallengeDescription = `
Creates a short-lived TON Connect authentication payload.`
)

// TONChallenge creates a TON Connect payload for HTTP authentication.
var TONChallenge = adapter.Method[struct{}, TONChallengeResponse]{
	Key:         tonChallengeKey,
	Description: tonChallengeDescription,
	Transports:  adapter.HTTP,
	Method:      fiber.MethodGet,
	Handler: func(
		_ *adapter.Context,
		_ struct{},
	) (TONChallengeResponse, error) {
		payload, err := controlauth.TONConnectPayload(
			config.ControlAuthTONPayloadSecret,
			config.ControlAuthTONMaxAge,
		)
		if err != nil {
			return TONChallengeResponse{}, err
		}

		return TONChallengeResponse{Payload: payload}, nil
	},
}
