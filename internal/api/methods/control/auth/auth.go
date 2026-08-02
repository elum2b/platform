package auth

import (
	"strings"
	"time"

	controlmodel "github.com/elum2b/services/control/model"
	controladmin "github.com/elum2b/services/control/service/admin"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/cookies"
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
	Account           *AccountResponse `json:"account"`
	SessionID         string           `json:"session_id"`
	SessionExpiresAt  time.Time        `json:"session_expires_at"`
}

func complete(
	ctx *adapter.Context,
	identity controladmin.AuthIdentityParams,
) (AuthResponse, error) {
	result, err := services.Control.Admin.CompleteAuth(ctx.Context, identity)
	if err != nil {
		return AuthResponse{}, err
	}

	return authResponse(ctx, result), nil
}

func authResponse(
	ctx *adapter.Context,
	result controladmin.AuthResult,
) AuthResponse {
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
			ctx.HTTP,
			config.ControlAuthTwoFactorCookieName,
			result.TwoFactorChallenge,
			time.Now().Add(config.ControlAuthTwoFactorDuration),
		)

		return response
	}

	cookies.Clear(ctx.HTTP, config.ControlAuthTwoFactorCookieName)
	cookies.Set(
		ctx.HTTP,
		config.ControlAuthCookieName,
		result.SessionToken,
		result.Session.ExpiresAt,
	)

	response.SessionID = result.Session.ID
	response.SessionExpiresAt = result.Session.ExpiresAt

	return response
}

func identityMetadata(
	ctx *adapter.Context,
	inviteToken string,
	bindToIP bool,
) controladmin.AuthIdentityParams {
	return controladmin.AuthIdentityParams{
		InviteToken: strings.TrimSpace(inviteToken),
		IP:          strings.TrimSpace(ctx.HTTP.IP()),
		UserAgent: strings.TrimSpace(
			ctx.HTTP.Get("User-Agent"),
		),
		BindToIP:  bindToIP,
		ExpiresAt: time.Now().Add(config.ControlAuthSessionDuration),
	}
}

func accountResponse(account controladmin.AccountModel) *AccountResponse {
	return &AccountResponse{
		ID:          account.ID,
		DisplayName: account.DisplayName,
		Status:      account.Status,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}
}
