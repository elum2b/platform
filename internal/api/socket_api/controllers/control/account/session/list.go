package session

import (
	"time"

	etp "github.com/elum-utils/go-etp"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type SessionListResponse struct {
	ID         string    `json:"id"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	BindToIP   bool      `json:"bind_to_ip"`
	ExpiresAt  time.Time `json:"expires_at"`
	LastUsedAt time.Time `json:"last_used_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func List(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		sessions, err := services.Control.Admin.ListSessions(
			ctx,
			ctx.Peer.Identity().UserID,
		)
		if err != nil {
			return err
		}

		response := make([]SessionListResponse, 0, len(sessions))
		for _, session := range sessions {
			response = append(response, SessionListResponse{
				ID:         session.ID,
				IP:         session.IP,
				UserAgent:  session.UserAgent,
				BindToIP:   session.BindToIP,
				ExpiresAt:  session.ExpiresAt,
				LastUsedAt: session.LastUsedAt,
				CreatedAt:  session.CreatedAt,
			})
		}

		return socketutils.Respond(ctx, event, response)
	})
}
