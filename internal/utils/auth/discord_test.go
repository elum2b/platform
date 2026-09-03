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
	})}

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
	})}

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

func discordRoundTripper(
	t *testing.T,
	bodies map[string]string,
) http.RoundTripper {
	t.Helper()

	return roundTripperFunc(
		func(request *http.Request) (*http.Response, error) {
			t.Helper()

			if request.Header.Get("Authorization") != "Bearer access-token" {
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
