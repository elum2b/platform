package offer

import (
	"time"

	cpaadmin "github.com/elum2b/services/cpa/service/admin"
	json "github.com/goccy/go-json"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID       string          `json:"workspace_id"                 validate:"required,uuid"`
	ID                string          `json:"id"                           validate:"required,max=255"`
	Payload           json.RawMessage `json:"payload"                      validate:"required"`
	Target            json.RawMessage `json:"target,omitempty"`
	CodeMode          string          `json:"code_mode"                    validate:"required"`
	CodeSource        *string         `json:"code_source,omitempty"`
	SharedCode        *string         `json:"shared_code,omitempty"`
	GeneratedLength   *int16          `json:"generated_length,omitempty"`
	GeneratedAlphabet *string         `json:"generated_alphabet,omitempty"`
	IsActive          bool            `json:"is_active"`
	StartAt           *time.Time      `json:"start_at,omitempty"`
	EndAt             *time.Time      `json:"end_at,omitempty"`
}

var (
	upsertKey         = "cpa.offer.upsert"
	upsertDescription = `
Creates or updates a CPA offer. Requires the 'cpa.offer.upsert' permission in
the target workspace.`
)

// Upsert exposes the CPA offer upsert method.
var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(upsertKey),
	},
	Handler: func(ctx *adapter.Context, data UpsertRequest) (struct{}, error) {
		err := services.CPA.Admin.UpsertOffer(
			ctx.Context,
			cpaadmin.UpsertOfferParams{
				WorkspaceID:       data.WorkspaceID,
				ID:                data.ID,
				Payload:           data.Payload,
				Target:            data.Target,
				CodeMode:          data.CodeMode,
				CodeSource:        data.CodeSource,
				SharedCode:        data.SharedCode,
				GeneratedLength:   data.GeneratedLength,
				GeneratedAlphabet: data.GeneratedAlphabet,
				IsActive:          data.IsActive,
				StartAt:           data.StartAt,
				EndAt:             data.EndAt,
			},
		)

		return struct{}{}, err
	},
}
