package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"

	"github.com/elum2b/platform/internal/config"
)

var (
	googleKey         = "control.auth.google"
	googleDescription = `
Authenticates an account through Google OAuth.`
)

// Google authenticates an account through Google OAuth.
var Google = oauthMethod(
	googleKey,
	googleDescription,
	oauthConfig{
		ClientID:     config.ControlAuthGoogleClientID,
		ClientSecret: config.ControlAuthGoogleClientSecret,
		TokenURL:     config.ControlAuthGoogleTokenURL,
		UserInfoURL:  config.ControlAuthGoogleUserInfoURL,
	},
	controlauth.Google,
)
