package auth

import (
	controlauth "github.com/elum2b/services/control/adapters"

	"github.com/elum2b/platform/internal/config"
)

var (
	gitLabKey         = "control.auth.gitlab"
	gitLabDescription = `
Authenticates an account through GitLab OAuth.`
)

// GitLab authenticates an account through GitLab OAuth.
var GitLab = oauthMethod(
	gitLabKey,
	gitLabDescription,
	oauthConfig{
		ClientID:     config.ControlAuthGitLabClientID,
		ClientSecret: config.ControlAuthGitLabClientSecret,
		TokenURL:     config.ControlAuthGitLabTokenURL,
		UserInfoURL:  config.ControlAuthGitLabUserInfoURL,
	},
	controlauth.GitLab,
)
