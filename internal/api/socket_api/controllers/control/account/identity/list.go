package identity

import (
	"time"

	etp "github.com/elum-utils/go-etp"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type IdentityListResponse struct {
	AccountID string    `json:"account_id"`
	Provider  string    `json:"provider"`
	Subject   string    `json:"subject"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func List(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		identities, err := services.Control.Admin.ListIdentities(
			ctx,
			ctx.Peer.Identity().UserID,
		)
		if err != nil {
			return err
		}

		response := make([]IdentityListResponse, 0, len(identities))
		for _, identity := range identities {
			response = append(response, IdentityListResponse{
				AccountID: identity.AccountID,
				Provider:  identity.Provider,
				Subject:   identity.Subject,
				CreatedAt: identity.CreatedAt,
				UpdatedAt: identity.UpdatedAt,
			})
		}

		return socketutils.Respond(ctx, event, response)
	})
}
