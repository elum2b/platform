package auth

import (
	"github.com/elum2b/platform/internal/config"
	authutils "github.com/elum2b/platform/internal/utils/auth"
)

var (
	discordKey         = "control.auth.discord"
	discordDescription = `
Authenticates an account through a Discord OAuth access token.`
)

// Discord authenticates an account through a Discord OAuth access token.
var Discord = oauthMethod(
	discordKey,
	discordDescription,
	oauthConfig{ClientID: config.ControlAuthDiscordClientID},
	authutils.Discord,
)
