package user

import (
	"time"

	puser "github.com/elum2b/services/payment/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type CreateAttemptRequest struct {
	WorkspaceID            string     `json:"workspace_id"                       query:"workspace_id"             validate:"required,uuid"`
	AppID                  int64      `json:"app_id"                             query:"app_id"                   validate:"required,min=1"`
	PlatformID             int64      `json:"platform_id"                        query:"platform_id"              validate:"required,min=1"`
	OrderID                uint64     `json:"order_id"                           query:"order_id"                 validate:"required,min=1"`
	ProviderCode           string     `json:"provider_code"                      query:"provider_code"            validate:"required,max=255"`
	ProviderPaymentID      *string    `json:"provider_payment_id,omitempty"      query:"provider_payment_id"`
	ProviderInvoiceID      *string    `json:"provider_invoice_id,omitempty"      query:"provider_invoice_id"`
	ProviderChargeID       *string    `json:"provider_charge_id,omitempty"       query:"provider_charge_id"`
	ProviderSubscriptionID *string    `json:"provider_subscription_id,omitempty" query:"provider_subscription_id"`
	IdempotencyKey         *string    `json:"idempotency_key,omitempty"          query:"idempotency_key"`
	ConfirmationURL        *string    `json:"confirmation_url,omitempty"         query:"confirmation_url"`
	ReturnURL              *string    `json:"return_url,omitempty"               query:"return_url"`
	ExpiresAt              *time.Time `json:"expires_at,omitempty"               query:"expires_at"`
}

type CreateAttemptResponse struct {
	Attempt puser.AttemptModel `json:"attempt"`
}

var (
	createAttemptKey         = "payment.user.create_attempt"
	createAttemptDescription = `
Creates a payment attempt for the authenticated application user.`
)

var CreateAttempt = adapter.Method[CreateAttemptRequest, CreateAttemptResponse]{
	Key:         createAttemptKey,
	Description: createAttemptDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, d CreateAttemptRequest) (CreateAttemptResponse, error) {
		v, err := services.Payment.User.CreateAttempt(
			ctx.Context,
			puser.CreateAttemptParams{
				Identity:               *ctx.Identity,
				OrderID:                d.OrderID,
				ProviderCode:           d.ProviderCode,
				ProviderPaymentID:      d.ProviderPaymentID,
				ProviderInvoiceID:      d.ProviderInvoiceID,
				ProviderChargeID:       d.ProviderChargeID,
				ProviderSubscriptionID: d.ProviderSubscriptionID,
				IdempotencyKey:         d.IdempotencyKey,
				ConfirmationURL:        d.ConfirmationURL,
				ReturnURL:              d.ReturnURL,
				ExpiresAt:              d.ExpiresAt,
			},
		)
		if err != nil {
			return CreateAttemptResponse{}, err
		}

		return CreateAttemptResponse{Attempt: *v}, nil
	},
}
