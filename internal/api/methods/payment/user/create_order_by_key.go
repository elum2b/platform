package user

import (
	"time"

	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateOrderByKeyRequest struct {
	WorkspaceID   string     `json:"workspace_id"             query:"workspace_id"   validate:"required,uuid"`
	AppID         int64      `json:"app_id"                   query:"app_id"         validate:"required,min=1"`
	PlatformID    int64      `json:"platform_id"              query:"platform_id"    validate:"required,min=1"`
	Key           string     `json:"key"                      query:"key"            validate:"required,max=255"`
	Payer         string     `json:"payer"                    query:"payer"          validate:"required"`
	AssetCode     string     `json:"asset_code"               query:"asset_code"     validate:"required,max=255"`
	Quantity      uint64     `json:"quantity"                 query:"quantity"       validate:"required,min=1"`
	Locale        string     `json:"locale,omitempty"         query:"locale"`
	ReservedUntil *time.Time `json:"reserved_until,omitempty" query:"reserved_until"`
	ExpiresAt     *time.Time `json:"expires_at,omitempty"     query:"expires_at"`
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
	Transports:  adapter.HTTP,
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
