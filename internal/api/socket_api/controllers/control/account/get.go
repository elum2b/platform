package account

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	controlmodel "github.com/elum2b/services/control/model"

	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type GetResponse struct {
	ID          string                     `json:"id"`
	DisplayName string                     `json:"display_name"`
	Status      controlmodel.AccountStatus `json:"status"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

func Get(event string, socket etp.Router) {
	socket.On(event, func(ctx *etp.Context) error {
		account, err := services.Control.Admin.GetAccount(
			ctx,
			ctx.Peer.Identity().UserID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, GetResponse{
			ID:          account.ID,
			DisplayName: account.DisplayName,
			Status:      account.Status,
			CreatedAt:   account.CreatedAt,
			UpdatedAt:   account.UpdatedAt,
		})
	})
}
