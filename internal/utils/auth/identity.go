package auth

import (
	"context"
	"strings"

	controlauth "github.com/elum2b/services/control/adapters"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/config"
)

type IdentityRequest struct {
	Provider        string                      `json:"provider"          validate:"required"`
	Code            string                      `json:"code"`
	AccessToken     string                      `json:"access_token"`
	RedirectURI     string                      `json:"redirect_uri"`
	InitData        string                      `json:"init_data"`
	Address         string                      `json:"address"`
	Network         string                      `json:"network"`
	PublicKey       string                      `json:"public_key"`
	WalletStateInit string                      `json:"wallet_state_init"`
	Proof           controlauth.TONConnectProof `json:"proof"`
}

type oauthResolver func(context.Context, controlauth.OAuth2AuthParams) (admin.AuthIdentityParams, error)

func ResolveIdentity(
	ctx context.Context,
	data IdentityRequest,
) (admin.AuthIdentityParams, error) {
	switch strings.ToLower(strings.TrimSpace(data.Provider)) {
	case controlauth.ProviderVKID:
		return resolveOAuth(
			ctx,
			data,
			config.ControlAuthVKClientID,
			config.ControlAuthVKClientSecret,
			config.ControlAuthVKTokenURL,
			config.ControlAuthVKUserInfoURL,
			controlauth.VKID,
		)
	case controlauth.ProviderGitHub:
		return resolveOAuth(
			ctx,
			data,
			config.ControlAuthGitHubClientID,
			config.ControlAuthGitHubClientSecret,
			config.ControlAuthGitHubTokenURL,
			config.ControlAuthGitHubUserInfoURL,
			controlauth.GitHub,
		)
	case controlauth.ProviderGitLab:
		return resolveOAuth(
			ctx,
			data,
			config.ControlAuthGitLabClientID,
			config.ControlAuthGitLabClientSecret,
			config.ControlAuthGitLabTokenURL,
			config.ControlAuthGitLabUserInfoURL,
			controlauth.GitLab,
		)
	case controlauth.ProviderGoogle:
		return resolveOAuth(
			ctx,
			data,
			config.ControlAuthGoogleClientID,
			config.ControlAuthGoogleClientSecret,
			config.ControlAuthGoogleTokenURL,
			config.ControlAuthGoogleUserInfoURL,
			controlauth.Google,
		)
	case controlauth.ProviderYandex:
		return resolveOAuth(
			ctx,
			data,
			config.ControlAuthYandexClientID,
			config.ControlAuthYandexClientSecret,
			config.ControlAuthYandexTokenURL,
			config.ControlAuthYandexUserInfoURL,
			controlauth.Yandex,
		)
	case controlauth.ProviderTelegramWebApp:
		return controlauth.TelegramWebApp(
			ctx,
			controlauth.TelegramWebAppAuthParams{
				BotToken: config.ControlAuthTelegramBotToken,
				InitData: data.InitData,
				MaxAge:   config.ControlAuthTelegramMaxAge,
			},
		)
	case controlauth.ProviderTONConnect:
		return controlauth.TONConnect(ctx, controlauth.TONConnectAuthParams{
			Address:           data.Address,
			Network:           data.Network,
			PublicKey:         data.PublicKey,
			WalletStateInit:   data.WalletStateInit,
			Proof:             data.Proof,
			PayloadSecret:     config.ControlAuthTONPayloadSecret,
			ExpectedDomain:    config.ControlAuthTONDomain,
			ExpectedNetwork:   config.ControlAuthTONNetwork,
			AllowNativeDomain: config.ControlAuthTONAllowNativeDomain,
			MaxAge:            config.ControlAuthTONMaxAge,
		})
	default:
		return admin.AuthIdentityParams{}, serviceerrors.ErrInvalidFields
	}
}

func resolveOAuth(
	ctx context.Context,
	data IdentityRequest,
	clientID string,
	clientSecret string,
	tokenURL string,
	userInfoURL string,
	resolve oauthResolver,
) (admin.AuthIdentityParams, error) {
	return resolve(ctx, controlauth.OAuth2AuthParams{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         data.Code,
		AccessToken:  data.AccessToken,
		RedirectURI:  data.RedirectURI,
		TokenURL:     tokenURL,
		UserInfoURL:  userInfoURL,
		Timeout:      config.ControlAuthOAuthTimeout,
	})
}
