package auth

import (
	"strings"

	"github.com/elum2b/platform/internal/config"
)

// Provider contains the public configuration needed to start authentication.
type Provider struct {
	Key      string `json:"key"`
	ClientID string `json:"client_id,omitempty"`
}

// Providers returns authentication providers fully configured for use.
func Providers() []Provider {
	providers := make([]Provider, 0, 8)

	addOAuthProvider(
		&providers,
		ProviderDiscord,
		config.ControlAuthDiscordClientID,
	)
	addOAuthProvider(&providers, "vk", config.ControlAuthVKClientID,
		config.ControlAuthVKClientSecret)
	addOAuthProvider(&providers, "github", config.ControlAuthGitHubClientID,
		config.ControlAuthGitHubClientSecret)
	addOAuthProvider(&providers, "gitlab", config.ControlAuthGitLabClientID,
		config.ControlAuthGitLabClientSecret)
	addOAuthProvider(&providers, "google", config.ControlAuthGoogleClientID,
		config.ControlAuthGoogleClientSecret)
	addOAuthProvider(&providers, "yandex", config.ControlAuthYandexClientID,
		config.ControlAuthYandexClientSecret)

	if configured(config.ControlAuthTelegramBotToken) {
		providers = append(providers, Provider{Key: "telegram"})
	}

	if configured(
		config.ControlAuthTONPayloadSecret,
		config.ControlAuthTONDomain,
	) {
		providers = append(providers, Provider{Key: "ton"})
	}

	return providers
}

func addOAuthProvider(providers *[]Provider, key string, values ...string) {
	if !configured(values...) {
		return
	}

	*providers = append(*providers, Provider{
		Key:      key,
		ClientID: strings.TrimSpace(values[0]),
	})
}

func configured(values ...string) bool {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}

	return true
}
