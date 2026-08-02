package identity

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListResponse struct {
	AccountID string    `json:"account_id"`
	Provider  string    `json:"provider"`
	Subject   string    `json:"subject"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	listKey         = "control.account.identity.list"
	listDescription = `
Lists identities bound to the authenticated account.`
)

// List returns identities bound to the authenticated account.
var List = adapter.Method[struct{}, []ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, _ struct{}) ([]ListResponse, error) {
		identities, err := services.Control.Admin.ListIdentities(
			ctx.Context,
			ctx.AccountID,
		)
		if err != nil {
			return nil, err
		}

		response := make([]ListResponse, 0, len(identities))
		for _, identity := range identities {
			response = append(response, ListResponse{
				AccountID: identity.AccountID,
				Provider:  identity.Provider,
				Subject:   identity.Subject,
				CreatedAt: identity.CreatedAt,
				UpdatedAt: identity.UpdatedAt,
			})
		}

		return response, nil
	},
}
