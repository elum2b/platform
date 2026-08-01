package socket

import (
	"testing"

	etp "github.com/elum-utils/go-etp"
)

type decodeRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name: "valid",
			body: `{"workspace_id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name:    "malformed json",
			body:    `{`,
			wantErr: true,
		},
		{
			name:    "invalid workspace id",
			body:    `{"workspace_id":"workspace"}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := &etp.Context{
				Body: etp.NewBytesBody([]byte(test.body)),
			}

			request := new(decodeRequest)
			valid := Decode(ctx, request)

			if valid != !test.wantErr {
				t.Fatalf("Decode() = %t, want %t", valid, !test.wantErr)
			}
		})
	}
}
