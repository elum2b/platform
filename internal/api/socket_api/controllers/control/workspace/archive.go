package workspace

import (
	etp "github.com/elum-utils/go-etp"
	serviceerrors "github.com/elum2b/services/errors"

	"github.com/elum2b/platform/internal/api/socket_api/middleware"
	"github.com/elum2b/platform/internal/services"
	socketutils "github.com/elum2b/platform/internal/utils/socket"
)

type ArchiveRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

type ArchiveResponse struct {
	Affected int64 `json:"affected"`
}

func Archive(event string, socket etp.Router) {
	socket.Use(event, middleware.WorkspaceAccess("control.workspace.archive"))

	socket.On(event, func(ctx *etp.Context) error {
		data := new(ArchiveRequest)
		if !socketutils.Decode(ctx, data) {
			return serviceerrors.ErrInvalidFields
		}

		affected, err := services.Control.Admin.ArchiveWorkspace(
			ctx,
			ctx.Peer.Identity().UserID,
			data.WorkspaceID,
		)
		if err != nil {
			return err
		}

		return socketutils.Respond(
			ctx,
			event,
			ArchiveResponse{Affected: affected},
		)
	})
}
