package promo

import (
	"encoding/json"
	"time"

	promoadmin "github.com/elum2b/services/promo/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID    string          `json:"workspace_id"              validate:"required,uuid"`
	ID             uint64          `json:"id,omitempty"`
	Code           string          `json:"code"                      validate:"required"`
	Payload        json.RawMessage `json:"payload"                   validate:"required"`
	Target         json.RawMessage `json:"target,omitempty"`
	MaxActivations uint64          `json:"max_activations"`
	IsActive       bool            `json:"is_active"`
	StartAt        *time.Time      `json:"start_at,omitempty"`
	EndAt          *time.Time      `json:"end_at,omitempty"`
}

type UpsertResponse struct {
	ID       uint64 `json:"id,omitempty"`
	Affected int64  `json:"affected,omitempty"`
}

var (
	upsertKey         = "promo.upsert"
	upsertDescription = `
Creates or updates a promo code. Requires the 'promo.upsert' permission in the
target workspace.`
)

// Upsert exposes the promo upsert method.
var Upsert = adapter.Method[UpsertRequest, UpsertResponse]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (UpsertResponse, error) {
		params := promoadmin.SavePromoParams{
			WorkspaceID:    data.WorkspaceID,
			ID:             data.ID,
			Code:           data.Code,
			Payload:        data.Payload,
			Target:         data.Target,
			MaxActivations: data.MaxActivations,
			IsActive:       data.IsActive,
			StartAt:        data.StartAt,
			EndAt:          data.EndAt,
		}

		if data.ID == 0 {
			id, err := services.Promo.Admin.CreatePromo(ctx.Context, params)

			return UpsertResponse{ID: id}, err
		}

		affected, err := services.Promo.Admin.UpdatePromo(ctx.Context, params)

		return UpsertResponse{Affected: affected}, err
	},
}
