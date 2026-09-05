package auth

import (
	"reflect"
	"testing"

	"github.com/elum2b/platform/internal/config"
)

func TestProvidersReturnsFullyConfiguredProviders(t *testing.T) {
	settings := []*string{
		&config.ControlAuthDiscordClientID,
		&config.ControlAuthDiscordClientSecret,
		&config.ControlAuthDiscordRedirectURI,
		&config.ControlAuthVKClientID,
		&config.ControlAuthVKClientSecret,
		&config.ControlAuthVKRedirectURI,
		&config.ControlAuthGitHubClientID,
		&config.ControlAuthGitHubClientSecret,
		&config.ControlAuthGitHubRedirectURI,
		&config.ControlAuthGitLabClientID,
		&config.ControlAuthGitLabClientSecret,
		&config.ControlAuthGitLabRedirectURI,
		&config.ControlAuthGoogleClientID,
		&config.ControlAuthGoogleClientSecret,
		&config.ControlAuthGoogleRedirectURI,
		&config.ControlAuthYandexClientID,
		&config.ControlAuthYandexClientSecret,
		&config.ControlAuthYandexRedirectURI,
		&config.ControlAuthTelegramBotToken,
		&config.ControlAuthTONPayloadSecret,
		&config.ControlAuthTONDomain,
	}
	saved := make([]string, len(settings))

	for index, setting := range settings {
		saved[index] = *setting
		*setting = ""
	}

	t.Cleanup(func() {
		for index, setting := range settings {
			*setting = saved[index]
		}
	})

	config.ControlAuthDiscordClientID = "discord-client"
	config.ControlAuthDiscordClientSecret = "discord-secret"
	config.ControlAuthDiscordRedirectURI = "https://app.example.com/discord"
	config.ControlAuthVKClientID = "vk-client"
	config.ControlAuthVKClientSecret = "vk-secret"
	config.ControlAuthVKRedirectURI = "https://app.example.com/vk"
	config.ControlAuthGitHubClientID = "github-client"
	config.ControlAuthGitHubClientSecret = "github-secret"
	config.ControlAuthGitHubRedirectURI = "https://app.example.com/github"
	config.ControlAuthGitLabClientID = "gitlab-client"
	config.ControlAuthGitLabClientSecret = "gitlab-secret"
	config.ControlAuthGitLabRedirectURI = "https://app.example.com/gitlab"
	config.ControlAuthGoogleClientID = "google-client"
	config.ControlAuthGoogleClientSecret = "google-secret"
	config.ControlAuthGoogleRedirectURI = "https://app.example.com/google"
	config.ControlAuthYandexClientID = "yandex-client"
	config.ControlAuthYandexClientSecret = "yandex-secret"
	config.ControlAuthYandexRedirectURI = "https://app.example.com/yandex"
	config.ControlAuthTelegramBotToken = "telegram-token"
	config.ControlAuthTONPayloadSecret = "ton-secret"
	config.ControlAuthTONDomain = "core.elumapp.ru"

	want := []Provider{
		{Key: "discord", ClientID: "discord-client"},
		{Key: "vk", ClientID: "vk-client"},
		{Key: "github", ClientID: "github-client"},
		{Key: "gitlab", ClientID: "gitlab-client"},
		{Key: "google", ClientID: "google-client"},
		{Key: "yandex", ClientID: "yandex-client"},
		{Key: "telegram"},
		{Key: "ton"},
	}
	if got := Providers(); !reflect.DeepEqual(got, want) {
		t.Errorf("Providers() = %#v, want %#v", got, want)
	}
}

func TestProvidersSkipsIncompleteConfiguration(t *testing.T) {
	settings := []*string{
		&config.ControlAuthDiscordClientID,
		&config.ControlAuthDiscordClientSecret,
		&config.ControlAuthDiscordRedirectURI,
		&config.ControlAuthVKClientID,
		&config.ControlAuthVKClientSecret,
		&config.ControlAuthVKRedirectURI,
		&config.ControlAuthGitHubClientID,
		&config.ControlAuthGitHubClientSecret,
		&config.ControlAuthGitHubRedirectURI,
		&config.ControlAuthGitLabClientID,
		&config.ControlAuthGitLabClientSecret,
		&config.ControlAuthGitLabRedirectURI,
		&config.ControlAuthGoogleClientID,
		&config.ControlAuthGoogleClientSecret,
		&config.ControlAuthGoogleRedirectURI,
		&config.ControlAuthYandexClientID,
		&config.ControlAuthYandexClientSecret,
		&config.ControlAuthYandexRedirectURI,
		&config.ControlAuthTelegramBotToken,
		&config.ControlAuthTONPayloadSecret,
		&config.ControlAuthTONDomain,
	}
	saved := make([]string, len(settings))

	for index, setting := range settings {
		saved[index] = *setting
		*setting = ""
	}

	t.Cleanup(func() {
		for index, setting := range settings {
			*setting = saved[index]
		}
	})

	config.ControlAuthDiscordClientID = "discord-client"
	config.ControlAuthDiscordClientSecret = "discord-secret"
	config.ControlAuthDiscordRedirectURI = "https://app.example.com/discord"
	config.ControlAuthVKClientID = "vk-client"
	config.ControlAuthTONPayloadSecret = "ton-secret"

	want := []Provider{{Key: "discord", ClientID: "discord-client"}}
	if got := Providers(); !reflect.DeepEqual(got, want) {
		t.Errorf("Providers() = %#v, want %#v", got, want)
	}
}
