package auth

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	controlauth "github.com/elum2b/services/control/adapters"
	serviceerrors "github.com/elum2b/services/errors"
)

func TestDiscordResolvesTokenIdentity(t *testing.T) {
	client := &http.Client{Transport: discordRoundTripper(t, map[string]string{
		"/api/v10/oauth2/@me": `{
			"application":{"id":"application-1"},
			"user":{"id":"user-1"}
		}`,
		"/api/v10/users/@me": `{
			"id":"user-1",
			"username":"discord-user",
			"global_name":"Discord User"
		}`,
	}, "access-token")}

	identity, err := Discord(context.Background(), controlauth.OAuth2AuthParams{
		ClientID:    "application-1",
		AccessToken: "access-token",
		InviteToken: "invite-token",
		IP:          "127.0.0.1",
		UserAgent:   "platform-test",
		BindToIP:    true,
		ExpiresAt:   time.Date(2026, 9, 2, 0, 0, 0, 0, time.UTC),
		HTTPClient:  client,
	})
	if err != nil {
		t.Fatalf("Discord() error = %v", err)
	}

	if identity.Provider != ProviderDiscord {
		t.Errorf("Provider = %q, want %q", identity.Provider, ProviderDiscord)
	}

	if identity.Subject != "user-1" {
		t.Errorf("Subject = %q, want %q", identity.Subject, "user-1")
	}

	if identity.DisplayName != "Discord User" {
		t.Errorf(
			"DisplayName = %q, want %q",
			identity.DisplayName,
			"Discord User",
		)
	}
}

func TestDiscordRejectsTokenFromAnotherApplication(t *testing.T) {
	client := &http.Client{Transport: discordRoundTripper(t, map[string]string{
		"/api/v10/oauth2/@me": `{
			"application":{"id":"another-application"},
			"user":{"id":"user-1"}
		}`,
	}, "access-token")}

	_, err := Discord(context.Background(), controlauth.OAuth2AuthParams{
		ClientID:    "application-1",
		AccessToken: "access-token",
		HTTPClient:  client,
	})
	if serviceerrors.CodeOf(err) != serviceerrors.CodeUnauthorized {
		t.Fatalf(
			"error code = %q, want %q",
			serviceerrors.CodeOf(err),
			serviceerrors.CodeUnauthorized,
		)
	}
}

func TestDiscordExchangesAuthorizationCode(t *testing.T) {
	client := &http.Client{Transport: discordRoundTripper(t, map[string]string{
		"/api/v10/oauth2/token": `{"access_token":"access-token"}`,
		"/api/v10/oauth2/@me": `{
			"application":{"id":"application-1"},
			"user":{"id":"user-1"}
		}`,
		"/api/v10/users/@me": `{
			"id":"user-1",
			"username":"discord-user"
		}`,
	}, "access-token")}

	_, err := Discord(context.Background(), controlauth.OAuth2AuthParams{
		ClientID:     "application-1",
		ClientSecret: "client-secret",
		Code:         "authorization-code",
		RedirectURI:  "https://app.example.com/auth/discord",
		HTTPClient:   client,
	})
	if err != nil {
		t.Fatalf("Discord() error = %v", err)
	}
}

func discordRoundTripper(
	t *testing.T,
	bodies map[string]string,
	accessToken string,
) http.RoundTripper {
	t.Helper()

	return roundTripperFunc(
		func(request *http.Request) (*http.Response, error) {
			t.Helper()

			if request.URL.Path == "/api/v10/oauth2/token" {
				if err := request.ParseForm(); err != nil {
					t.Fatalf("ParseForm() error = %v", err)
				}

				if request.Form.Get("grant_type") != "authorization_code" {
					t.Errorf("grant_type = %q", request.Form.Get("grant_type"))
				}
			} else if request.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Errorf(
					"Authorization = %q",
					request.Header.Get("Authorization"),
				)
			}

			body, ok := bodies[request.URL.Path]
			if !ok {
				t.Fatalf("unexpected request path %q", request.URL.Path)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
				Request:    request,
			}, nil
		},
	)
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (function roundTripperFunc) RoundTrip(
	request *http.Request,
) (*http.Response, error) {
	return function(request)
}
