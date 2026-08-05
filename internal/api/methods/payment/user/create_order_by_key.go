package user

import (
	"time"

	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateOrderByKeyRequest struct {
	WorkspaceID   string     `json:"workspace_id"   validate:"required,uuid"`
	AppID         int64      `json:"app_id"         validate:"required,min=1"`
	PlatformID    int64      `json:"platform_id"    validate:"required,min=1"`
	Params        string     `json:"params"         validate:"required"`
	Key           string     `json:"key"            validate:"required,max=255"`
	Payer         string     `json:"payer"          validate:"required"`
	AssetCode     string     `json:"asset_code"     validate:"required,max=255"`
	Quantity      uint64     `json:"quantity"       validate:"required,min=1"`
	Locale        string     `json:"locale,omitempty"`
	ReservedUntil *time.Time `json:"reserved_until,omitempty"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"`
}

type CreateOrderByKeyResponse struct {
	Order puser.OrderModel `json:"order"`
}

var (
	createOrderByKeyKey         = "payment.user.create_order_by_key"
	createOrderByKeyDescription = `
Creates an order by product key for the authenticated application user.`
)

var CreateOrderByKey = adapter.Method[CreateOrderByKeyRequest, CreateOrderByKeyResponse]{
	Key:         createOrderByKeyKey,
	Description: createOrderByKeyDescription,
	Transports:  adapter.WS,
	Handler: func(ctx *adapter.Context, d CreateOrderByKeyRequest) (CreateOrderByKeyResponse, error) {
		v, err := services.Payment.User.CreateOrderByKey(
			ctx.Context,
			puser.CreateOrderByKeyParams{
				Key:           d.Key,
				Payer:         makeActor(d.PlatformID, d.Payer, nil),
				AssetCode:     d.AssetCode,
				Quantity:      d.Quantity,
				Locale:        d.Locale,
				ReservedUntil: d.ReservedUntil,
				ExpiresAt:     d.ExpiresAt,
			},
		)
		if err != nil {
			return CreateOrderByKeyResponse{}, err
		}
		return CreateOrderByKeyResponse{Order: *v}, nil
	},
}
