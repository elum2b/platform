package promo

import (
	"github.com/elum2b/services/delivery"
	service "github.com/elum2b/services/promo"

	"github.com/elum2b/platform/internal/services"
)

func handler(ctx service.Context) error {
	return services.Delivery.DeliverCallback(ctx, ctx, delivery.Message{
		Destination: delivery.Destination{
			WorkspaceID: ctx.Payload.WorkspaceID,
			AppID:       ctx.Payload.AppID,
			PlatformID:  ctx.Payload.PlatformID,
		},
		EventType:      ctx.EventType,
		IdempotencyKey: ctx.IdempotencyKey,
		Payload:        ctx.Context.Payload,
	})
}
