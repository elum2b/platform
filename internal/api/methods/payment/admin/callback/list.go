package callback

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ListRequest struct {
	WorkspaceID   string `json:"workspace_id"           validate:"required,uuid"`
	SourceService string `json:"source_service,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	Status        string `json:"status,omitempty"`
	Limit         int32  `json:"limit,omitempty"        validate:"omitempty,min=1,max=100"`
	Offset        int32  `json:"offset,omitempty"       validate:"min=0"`
}

type ListEvent struct {
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

type ListResponse struct {
	Events []ListEvent `json:"events"`
}

var (
	listKey         = "payment.callback.list"
	listDescription = `
Lists payment callback delivery events and their states. Requires the
'payment.callback.list' permission in the target workspace.`
)

var List = adapter.Method[ListRequest, ListResponse]{
	Key:         listKey,
	Description: listDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(listKey)},
	Handler: func(ctx *adapter.Context, data ListRequest) (ListResponse, error) {
		values, err := services.Payment.Admin.ListCallbackEvents(
			ctx.Context,
			padm.CallbackEventListParams{
				WorkspaceID:   data.WorkspaceID,
				SourceService: data.SourceService,
				EventType:     data.EventType,
				Status:        data.Status,
				Page:          padm.PageParams{Limit: data.Limit, Offset: data.Offset},
			},
		)
		if err != nil {
			return ListResponse{}, err
		}

		events := make([]ListEvent, 0, len(values))

		for _, v := range values {
			events = append(
				events,
				ListEvent{
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
			)
		}

		return ListResponse{Events: events}, nil
	},
}
