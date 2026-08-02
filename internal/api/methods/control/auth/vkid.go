package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"

	"github.com/elum2b/platform/internal/config"
)

var (
	vkidKey         = "control.auth.vkid"
	vkidDescription = `
Authenticates an account through VK ID OAuth.`
)

// VKID authenticates an account through VK ID OAuth.
var VKID = oauthMethod(
	vkidKey,
	vkidDescription,
	oauthConfig{
		ClientID:     config.ControlAuthVKClientID,
		ClientSecret: config.ControlAuthVKClientSecret,
		TokenURL:     config.ControlAuthVKTokenURL,
		UserInfoURL:  config.ControlAuthVKUserInfoURL,
	},
	controlauth.VKID,
)
