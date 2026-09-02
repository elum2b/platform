package auth

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	controlauth "github.com/elum2b/services/control/adapters"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"
	json "github.com/goccy/go-json"
)

const (
	// ProviderDiscord identifies Discord OAuth identities.
	ProviderDiscord = "discord"

	discordAuthorizationURL = "https://discord.com/api/v10/oauth2/@me"
	discordUserURL          = "https://discord.com/api/v10/users/@me"
	discordResponseMaxSize  = 1 << 20
)

// Discord resolves a Discord identity from an OAuth access token.
func Discord(
	ctx context.Context,
	params controlauth.OAuth2AuthParams,
) (admin.AuthIdentityParams, error) {
	clientID := strings.TrimSpace(params.ClientID)
	if clientID == "" {
		return admin.AuthIdentityParams{}, controlauth.ErrClientIDRequired
	}

	token := strings.TrimSpace(params.AccessToken)
	if token == "" {
		return admin.AuthIdentityParams{}, controlauth.ErrTokenRequired
	}

	if params.Timeout > 0 {
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, params.Timeout)

		defer cancel()
	}

	client := params.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	authorization, err := discordAuthorization(ctx, client, token)
	if err != nil {
		return admin.AuthIdentityParams{}, err
	}

	if authorization.Application.ID != clientID {
		return admin.AuthIdentityParams{}, serviceerrors.New(
			serviceerrors.CodeUnauthorized,
			"discord token belongs to another application",
		)
	}

	user, raw, err := discordUser(ctx, client, token)
	if err != nil {
		return admin.AuthIdentityParams{}, err
	}

	if user.ID == "" {
		return admin.AuthIdentityParams{}, serviceerrors.New(
			serviceerrors.CodeUnauthorized,
			"discord user id is required",
		)
	}

	if authorization.User.ID != "" && authorization.User.ID != user.ID {
		return admin.AuthIdentityParams{}, serviceerrors.New(
			serviceerrors.CodeUnauthorized,
			"discord token user does not match",
		)
	}

	payload, err := json.Marshal(map[string]any{
		"profile": user,
		"raw":     json.RawMessage(raw),
	})
	if err != nil {
		return admin.AuthIdentityParams{}, err
	}

	return admin.AuthIdentityParams{
		Provider:    ProviderDiscord,
		Subject:     user.ID,
		DisplayName: discordDisplayName(user),
		Payload:     payload,
		InviteToken: strings.TrimSpace(params.InviteToken),
		IP:          strings.TrimSpace(params.IP),
		UserAgent:   strings.TrimSpace(params.UserAgent),
		BindToIP:    params.BindToIP,
		ExpiresAt:   params.ExpiresAt,
	}, nil
}

type discordAuthorizationResponse struct {
	Application struct {
		ID string `json:"id"`
	} `json:"application"`
	User struct {
		ID string `json:"id"`
	} `json:"user"`
}

type discordUserResponse struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	GlobalName string `json:"global_name"`
}

func discordAuthorization(
	ctx context.Context,
	client *http.Client,
	token string,
) (discordAuthorizationResponse, error) {
	response := new(discordAuthorizationResponse)
	if err := discordGet(ctx, client, token, discordAuthorizationURL, response); err != nil {
		return discordAuthorizationResponse{}, err
	}

	return *response, nil
}

func discordUser(
	ctx context.Context,
	client *http.Client,
	token string,
) (discordUserResponse, []byte, error) {
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		discordUserURL,
		nil,
	)
	if err != nil {
		return discordUserResponse{}, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return discordUserResponse{}, nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.LimitReader(response.Body, discordResponseMaxSize))
	if err != nil {
		return discordUserResponse{}, nil, err
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return discordUserResponse{}, nil, serviceerrors.New(
			serviceerrors.CodeUnauthorized,
			"discord user request failed",
		)
	}

	user := new(discordUserResponse)
	decoder := json.NewDecoder(bytes.NewReader(body))
	if err := decoder.Decode(user); err != nil {
		return discordUserResponse{}, nil, err
	}

	return *user, body, nil
}

func discordGet(
	ctx context.Context,
	client *http.Client,
	token string,
	url string,
	target any,
) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.LimitReader(response.Body, discordResponseMaxSize))
	if err != nil {
		return err
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return serviceerrors.New(
			serviceerrors.CodeUnauthorized,
			"discord authorization request failed",
		)
	}

	return json.Unmarshal(body, target)
}

func discordDisplayName(user discordUserResponse) string {
	if name := strings.TrimSpace(user.GlobalName); name != "" {
		return name
	}

	return strings.TrimSpace(user.Username)
}
