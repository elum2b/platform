package account

import (
	"time"

	controlmodel "github.com/elum2b/services/control/model"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetResponse struct {
	ID          string                     `json:"id"`
	DisplayName string                     `json:"display_name"`
	Status      controlmodel.AccountStatus `json:"status"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

var (
	getKey         = "control.account.get"
	getDescription = `
Gets the authenticated control account.`
)

// Get returns the authenticated account.
var Get = adapter.Method[struct{}, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, _ struct{}) (GetResponse, error) {
		account, err := services.Control.Admin.GetAccount(
			ctx.Context,
			ctx.AccountID,
		)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			ID:          account.ID,
			DisplayName: account.DisplayName,
			Status:      account.Status,
			CreatedAt:   account.CreatedAt,
			UpdatedAt:   account.UpdatedAt,
		}, nil
	},
}
