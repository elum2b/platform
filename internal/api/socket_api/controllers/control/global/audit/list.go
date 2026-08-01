package audit

import (
	"time"

	etp "github.com/elum-utils/go-etp"
	"github.com/elum2b/services/control/service/admin"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ListRequest struct {
	Limit    int32     `json:"limit"     validate:"omitempty,min=1,max=100"`
	CursorAt time.Time `json:"cursor_at"`
	CursorID string    `json:"cursor_id"`
}

func List(event string, socket etp.Router) {
	socket.Use(event, middleware.GlobalAccess("control.global.audit.list"))
	socket.On(event, func(ctx *etp.Context) error {
		data := new(ListRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		items, err := services.Control.Admin.ListGlobalAudit(
			ctx,
			admin.Page{
				Limit:    data.Limit,
				CursorAt: data.CursorAt,
				CursorID: data.CursorID,
			},
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(ctx, event, items)
	})
}
