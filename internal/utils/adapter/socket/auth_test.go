package socket

import (
	"context"
	"testing"
)

func TestSessionTokenContext(t *testing.T) {
	ctx := withSessionToken(context.Background(), " session-token ")

	if token := sessionToken(ctx); token != "session-token" {
		t.Fatalf("unexpected session token: %q", token)
	}
}
