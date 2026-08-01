package auth

import (
	"context"

	controlauth "github.com/elum2b/services/control/adapters"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"

	"github.com/elum2b/platform/internal/config"
	httputils "github.com/elum2b/platform/internal/utils/http"
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

type oauthResolver func(context.Context, controlauth.OAuth2AuthParams) (admin.AuthIdentityParams, error)

func VK(path string, app fiber.Router) {
	oauth(path, app, oauthConfig{
		ClientID:     config.ControlAuthVKClientID,
		ClientSecret: config.ControlAuthVKClientSecret,
		TokenURL:     config.ControlAuthVKTokenURL,
		UserInfoURL:  config.ControlAuthVKUserInfoURL,
	}, controlauth.VKID)
}

func GitHub(path string, app fiber.Router) {
	oauth(path, app, oauthConfig{
		ClientID:     config.ControlAuthGitHubClientID,
		ClientSecret: config.ControlAuthGitHubClientSecret,
		TokenURL:     config.ControlAuthGitHubTokenURL,
		UserInfoURL:  config.ControlAuthGitHubUserInfoURL,
	}, controlauth.GitHub)
}

func GitLab(path string, app fiber.Router) {
	oauth(path, app, oauthConfig{
		ClientID:     config.ControlAuthGitLabClientID,
		ClientSecret: config.ControlAuthGitLabClientSecret,
		TokenURL:     config.ControlAuthGitLabTokenURL,
		UserInfoURL:  config.ControlAuthGitLabUserInfoURL,
	}, controlauth.GitLab)
}

func Google(path string, app fiber.Router) {
	oauth(path, app, oauthConfig{
		ClientID:     config.ControlAuthGoogleClientID,
		ClientSecret: config.ControlAuthGoogleClientSecret,
		TokenURL:     config.ControlAuthGoogleTokenURL,
		UserInfoURL:  config.ControlAuthGoogleUserInfoURL,
	}, controlauth.Google)
}

func Yandex(path string, app fiber.Router) {
	oauth(path, app, oauthConfig{
		ClientID:     config.ControlAuthYandexClientID,
		ClientSecret: config.ControlAuthYandexClientSecret,
		TokenURL:     config.ControlAuthYandexTokenURL,
		UserInfoURL:  config.ControlAuthYandexUserInfoURL,
	}, controlauth.Yandex)
}

func oauth(
	path string,
	app fiber.Router,
	provider oauthConfig,
	resolve oauthResolver,
) {
	app.Post(path, func(ctx fiber.Ctx) error {
		if !controlReady() {
			return httputils.Error(ctx, serviceerrors.ErrNotReady)
		}

		data := new(OAuthRequest)
		if !httputils.Decode(ctx, data) {
			return httputils.Error(ctx, serviceerrors.ErrInvalidFields)
		}

		metadata := identityMetadata(ctx, data.InviteToken, data.BindToIP)

		identity, err := resolve(ctx.Context(), controlauth.OAuth2AuthParams{
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
		})
		if err != nil {
			return httputils.Error(ctx, err)
		}

		return complete(ctx, identity)
	})
}
