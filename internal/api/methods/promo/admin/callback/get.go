package callback

import (
	"time"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type GetRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
	ID          uint64 `json:"id"           validate:"required,min=1"`
}

type GetEvent struct {
	ID                 uint64     `json:"id"`
	WorkspaceID        string     `json:"workspace_id"`
	SourceService      string     `json:"source_service"`
	EventType          string     `json:"event_type"`
	EventKey           string     `json:"event_key"`
	IdempotencyKey     string     `json:"idempotency_key"`
	Payload            []byte     `json:"payload"`
	PayloadContentType string     `json:"payload_content_type"`
	Status             string     `json:"status"`
	AttemptCount       uint32     `json:"attempt_count"`
	NextAttemptAt      time.Time  `json:"next_attempt_at"`
	LockedBy           *string    `json:"locked_by,omitempty"`
	LockedUntil        *time.Time `json:"locked_until,omitempty"`
	DeliveredAt        *time.Time `json:"delivered_at,omitempty"`
	RejectedAt         *time.Time `json:"rejected_at,omitempty"`
	LastError          *string    `json:"last_error,omitempty"`
	RejectReason       *string    `json:"reject_reason,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

type GetResponse struct {
	Event GetEvent `json:"event"`
}

var (
	getKey         = "promo.callback.get"
	getDescription = `
Returns a promo callback delivery event. Requires the 'promo.callback.get'
permission in the target workspace.`
)

// Get exposes the promo callback event retrieval method.
var Get = adapter.Method[GetRequest, GetResponse]{
	Key:         getKey,
	Description: getDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(getKey),
	},
	Handler: func(ctx *adapter.Context, data GetRequest) (GetResponse, error) {
		v, err := services.Promo.Admin.GetCallbackEvent(
			ctx.Context,
			data.WorkspaceID,
			data.ID,
		)

		return GetResponse{
			Event: GetEvent{
				ID:                 v.ID,
				WorkspaceID:        v.WorkspaceID,
				SourceService:      v.SourceService,
				EventType:          v.EventType,
				EventKey:           v.EventKey,
				IdempotencyKey:     v.IdempotencyKey,
				Payload:            v.Payload,
				PayloadContentType: v.PayloadContentType,
				Status:             v.Status,
				AttemptCount:       v.AttemptCount,
				NextAttemptAt:      v.NextAttemptAt,
				LockedBy:           v.LockedBy,
				LockedUntil:        v.LockedUntil,
				DeliveredAt:        v.DeliveredAt,
				RejectedAt:         v.RejectedAt,
				LastError:          v.LastError,
				RejectReason:       v.RejectReason,
				CreatedAt:          v.CreatedAt,
				UpdatedAt:          v.UpdatedAt,
			},
		}, err
	},
}
