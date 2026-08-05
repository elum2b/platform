package operation

import (
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateProductKeyRequest struct {
	WorkspaceID    string `json:"workspace_id"      validate:"required,uuid"`
	AppID          int64  `json:"app_id"            validate:"required,min=1"`
	PlatformID     int64  `json:"platform_id"       validate:"required,min=1"`
	PlatformUserID string `json:"platform_user_id"  validate:"required"`
	InternalUserID *int64 `json:"internal_user_id,omitempty"`
	ProductID      string `json:"product_id"        validate:"required,max=255"`
	MaxUses        int32  `json:"max_uses,omitempty"`
	ExpiresAt      *int64 `json:"expires_at,omitempty"`
}

type CreateProductKeyResponse struct {
	Key string `json:"key"`
}

var (
	createProductKeyKey         = "payment.operation.create_product_key"
	createProductKeyDescription = `
Creates a product key. Requires the 'payment.operation.create_product_key'
permission in the target workspace.`
)

var CreateProductKey = adapter.Method[CreateProductKeyRequest, CreateProductKeyResponse]{
	Key:         createProductKeyKey,
	Description: createProductKeyDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(createProductKeyKey)},
	Handler: func(ctx *adapter.Context, d CreateProductKeyRequest) (CreateProductKeyResponse, error) {
		k, err := services.Payment.Admin.CreateProductKey(
			ctx.Context,
			padm.CreateProductKeyParams{
				WorkspaceID:    d.WorkspaceID,
				AppID:          d.AppID,
				PlatformID:     d.PlatformID,
				PlatformUserID: d.PlatformUserID,
				InternalUserID: d.InternalUserID,
				ProductID:      d.ProductID,
				MaxUses:        d.MaxUses,
				ExpiresAt:      int64ToTimePtr(d.ExpiresAt),
			})
		return CreateProductKeyResponse{Key: k}, err
	},
}

func int64ToTimePtr(ts *int64) *time.Time {
	if ts == nil {
		return nil
	}
	t := time.Unix(*ts, 0)
	return &t
}
