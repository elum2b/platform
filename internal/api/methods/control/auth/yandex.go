package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"

	"github.com/elum2b/platform/internal/config"
)

var (
	yandexKey         = "control.auth.yandex"
	yandexDescription = `
Authenticates an account through Yandex OAuth.`
)

// Yandex authenticates an account through Yandex OAuth.
var Yandex = oauthMethod(
	yandexKey,
	yandexDescription,
	oauthConfig{
		ClientID:     config.ControlAuthYandexClientID,
		ClientSecret: config.ControlAuthYandexClientSecret,
		TokenURL:     config.ControlAuthYandexTokenURL,
		UserInfoURL:  config.ControlAuthYandexUserInfoURL,
	},
	controlauth.Yandex,
)
