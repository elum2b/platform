package auth

import (
	"github.com/elum2b/platform/internal/config"
	authutils "github.com/elum2b/platform/internal/utils/auth"
)

var (
	discordKey         = "control.auth.discord"
	discordDescription = `
Authenticates an account through a Discord OAuth authorization code.`
)

// Discord authenticates an account through a Discord OAuth authorization code.
var Discord = oauthMethod(
	discordKey,
	discordDescription,
	oauthConfig{
		ClientID:     config.ControlAuthDiscordClientID,
		ClientSecret: config.ControlAuthDiscordClientSecret,
		RedirectURI:  config.ControlAuthDiscordRedirectURI,
	},
	authutils.Discord,
)
