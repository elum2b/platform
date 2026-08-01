package auth

import (
	"strings"
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
	"github.com/elum2b/platform/internal/utils/cookies"
	httputils "github.com/elum2b/platform/internal/utils/http"
)

type AccountResponse struct {
	ID          string                     `json:"id"`
	DisplayName string                     `json:"display_name"`
	Status      controlmodel.AccountStatus `json:"status"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

type AuthResponse struct {
	Authenticated     bool             `json:"authenticated"`
	TwoFactorRequired bool             `json:"two_factor_required"`
	Created           bool             `json:"created"`
	Account           *AccountResponse `json:"account,omitempty"`
	SessionID         string           `json:"session_id,omitempty"`
	SessionExpiresAt  time.Time        `json:"session_expires_at,omitempty"`
}

func complete(ctx fiber.Ctx, identity admin.AuthIdentityParams) error {
	if !controlReady() {
		return httputils.Error(ctx, serviceerrors.ErrNotReady)
	}

	result, err := services.Control.Admin.CompleteAuth(ctx.Context(), identity)
	if err != nil {
		return httputils.Error(ctx, err)
	}

	return respondAuth(ctx, result)
}

func respondAuth(ctx fiber.Ctx, result admin.AuthResult) error {
	response := AuthResponse{
		Authenticated:     !result.TwoFactorRequired,
		TwoFactorRequired: result.TwoFactorRequired,
		Created:           result.Created,
	}

	if result.Account.ID != "" {
		response.Account = accountResponse(result.Account)
	}

	if result.TwoFactorRequired {
		cookies.Set(
			ctx,
			config.ControlAuthTwoFactorCookieName,
			result.TwoFactorChallenge,
			time.Now().Add(config.ControlAuthTwoFactorDuration),
		)

		return httputils.Respond(ctx, response)
	}

	cookies.Clear(ctx, config.ControlAuthTwoFactorCookieName)
	cookies.Set(
		ctx,
		config.ControlAuthCookieName,
		result.SessionToken,
		result.Session.ExpiresAt,
	)

	response.SessionID = result.Session.ID
	response.SessionExpiresAt = result.Session.ExpiresAt

	return httputils.Respond(ctx, response)
}

func identityMetadata(
	ctx fiber.Ctx,
	inviteToken string,
	bindToIP bool,
) admin.AuthIdentityParams {
	return admin.AuthIdentityParams{
		InviteToken: strings.TrimSpace(inviteToken),
		IP:          strings.TrimSpace(ctx.IP()),
		UserAgent:   strings.TrimSpace(ctx.Get(fiber.HeaderUserAgent)),
		BindToIP:    bindToIP,
		ExpiresAt:   time.Now().Add(config.ControlAuthSessionDuration),
	}
}

func accountResponse(account admin.AccountModel) *AccountResponse {
	return &AccountResponse{
		ID:          account.ID,
		DisplayName: account.DisplayName,
		Status:      account.Status,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}
}

func controlReady() bool {
	return services.Control != nil && services.Control.IsReady()
}
