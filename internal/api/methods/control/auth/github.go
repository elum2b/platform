package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"

	"github.com/elum2b/platform/internal/config"
)

var (
	githubKey         = "control.auth.github"
	githubDescription = `
Authenticates an account through GitHub OAuth.`
)

// GitHub authenticates an account through GitHub OAuth.
var GitHub = oauthMethod(
	githubKey,
	githubDescription,
	oauthConfig{
		ClientID:     config.ControlAuthGitHubClientID,
		ClientSecret: config.ControlAuthGitHubClientSecret,
		TokenURL:     config.ControlAuthGitHubTokenURL,
		UserInfoURL:  config.ControlAuthGitHubUserInfoURL,
	},
	controlauth.GitHub,
)
