package user

import (
	"time"

	services "github.com/elum2b/services"
	puser "github.com/elum2b/services/payment/service/user"

	appservices "github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateOrderRequest struct {
	WorkspaceID    string     `json:"workspace_id"               validate:"required,uuid"`
	AppID          int64      `json:"app_id"                     validate:"required,min=1"`
	PlatformID     int64      `json:"platform_id"                validate:"required,min=1"`
	Params         string     `json:"params"                     validate:"required"`
	InternalUserID *int64     `json:"internal_user_id,omitempty"`
	Payer          string     `json:"payer"                      validate:"required"`
	ProductID      string     `json:"product_id"                 validate:"required,max=255"`
	Quantity       uint64     `json:"quantity"                   validate:"required,min=1"`
	AssetCode      string     `json:"asset_code"                 validate:"required,max=255"`
	Locale         string     `json:"locale,omitempty"`
	ReservedUntil  *time.Time `json:"reserved_until,omitempty"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
}

type CreateOrderResponse struct {
	Order puser.OrderModel `json:"order"`
}

var (
	createOrderKey         = "payment.user.create_order"
	createOrderDescription = `
Creates an order for the authenticated application user.`
)

var CreateOrder = adapter.Method[CreateOrderRequest, CreateOrderResponse]{
	Key:         createOrderKey,
	Description: createOrderDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d CreateOrderRequest) (CreateOrderResponse, error) {
		v, err := appservices.Payment.User.CreateOrder(
			ctx.Context,
			puser.CreateOrderParams{
				Identity:       *ctx.Identity,
				InternalUserID: d.InternalUserID,
				Payer: makeActor(
					d.PlatformID,
					d.Payer,
					d.InternalUserID,
				),
				ProductID:     d.ProductID,
				Quantity:      d.Quantity,
				AssetCode:     d.AssetCode,
				Locale:        d.Locale,
				ReservedUntil: d.ReservedUntil,
				ExpiresAt:     d.ExpiresAt,
			},
		)
		if err != nil {
			return CreateOrderResponse{}, err
		}

		return CreateOrderResponse{Order: *v}, nil
	},
}

func makeActor(
	platformID int64,
	platformUserID string,
	internalUserID *int64,
) *services.Actor {
	return &services.Actor{
		PlatformID:     platformID,
		PlatformUserID: platformUserID,
		InternalUserID: internalUserID,
	}
}
