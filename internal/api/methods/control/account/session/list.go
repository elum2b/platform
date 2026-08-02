package session

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListResponse struct {
	ID         string    `json:"id"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	BindToIP   bool      `json:"bind_to_ip"`
	ExpiresAt  time.Time `json:"expires_at"`
	LastUsedAt time.Time `json:"last_used_at"`
	CreatedAt  time.Time `json:"created_at"`
}

var (
	listKey         = "control.account.session.list"
	listDescription = `
Lists sessions for the authenticated account.`
)

// List returns active sessions for the authenticated account.
var List = adapter.Method[struct{}, []ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, _ struct{}) ([]ListResponse, error) {
		sessions, err := services.Control.Admin.ListSessions(
			ctx.Context,
			ctx.AccountID,
		)
		if err != nil {
			return nil, err
		}

		response := make([]ListResponse, 0, len(sessions))
		for _, session := range sessions {
			response = append(response, ListResponse{
				ID:         session.ID,
				IP:         session.IP,
				UserAgent:  session.UserAgent,
				BindToIP:   session.BindToIP,
				ExpiresAt:  session.ExpiresAt,
				LastUsedAt: session.LastUsedAt,
				CreatedAt:  session.CreatedAt,
			})
		}

		return response, nil
	},
}
