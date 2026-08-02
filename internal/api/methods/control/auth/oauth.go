package auth

import (
	"context"

	controlauth "github.com/elum2b/services/control/adapters"
	controladmin "github.com/elum2b/services/control/service/admin"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type OAuthRequest struct {
	Code        string `json:"code"`
	AccessToken string `json:"access_token"`
	RedirectURI string `json:"redirect_uri"`
	InviteToken string `json:"invite_token"`
	BindToIP    bool   `json:"bind_to_ip"`
}

type oauthConfig struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	UserInfoURL  string
}

type oauthResolver func(
	context.Context,
	controlauth.OAuth2AuthParams,
) (controladmin.AuthIdentityParams, error)

func oauthMethod(
	key string,
	description string,
	provider oauthConfig,
	resolve oauthResolver,
) adapter.Method[OAuthRequest, AuthResponse] {
	return adapter.Method[OAuthRequest, AuthResponse]{
		Key:         key,
		Description: description,
		Transports:  adapter.HTTP,
		Method:      fiber.MethodPost,
		Handler: func(
			ctx *adapter.Context,
			data OAuthRequest,
		) (AuthResponse, error) {
			metadata := identityMetadata(ctx, data.InviteToken, data.BindToIP)

			identity, err := resolve(
				ctx.Context,
				controlauth.OAuth2AuthParams{
					ClientID:     provider.ClientID,
					ClientSecret: provider.ClientSecret,
					InviteToken:  metadata.InviteToken,
					Code:         data.Code,
					AccessToken:  data.AccessToken,
					RedirectURI:  data.RedirectURI,
					IP:           metadata.IP,
					UserAgent:    metadata.UserAgent,
					BindToIP:     metadata.BindToIP,
					ExpiresAt:    metadata.ExpiresAt,
					TokenURL:     provider.TokenURL,
					UserInfoURL:  provider.UserInfoURL,
					Timeout:      config.ControlAuthOAuthTimeout,
				},
			)
			if err != nil {
				return AuthResponse{}, err
			}

			return complete(ctx, identity)
		},
	}
}
