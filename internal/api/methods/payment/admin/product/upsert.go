package product

import (
	"encoding/json"
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID          string  `json:"workspace_id"                     validate:"required,uuid"`
	ID                   string  `json:"id,omitempty"`
	GroupCode            *string `json:"group_code,omitempty"`
	TitleKey             string  `json:"title_key"                        validate:"required,max=255"`
	DescriptionKey       *string `json:"description_key,omitempty"`
	Target               string  `json:"target"                           validate:"required"`
	ImageURL             *string `json:"image_url,omitempty"`
	LinkURL              *string `json:"link_url,omitempty"`
	SizeLabel            *string `json:"size_label,omitempty"`
	PeriodSeconds        *int64  `json:"period_seconds,omitempty"`
	TrialDurationSeconds *int64  `json:"trial_duration_seconds,omitempty"`
	QuantityMode         string  `json:"quantity_mode"                    validate:"required"`
	Position             int32   `json:"position"`
	GlobalLimit          int32   `json:"global_limit"`
	GlobalInterval       string  `json:"global_interval"`
	GlobalIntervalCount  int32   `json:"global_interval_count"`
	UserLimit            int32   `json:"user_limit"`
	UserInterval         string  `json:"user_interval"`
	UserIntervalCount    int32   `json:"user_interval_count"`
	AvailableFrom        *int64  `json:"available_from,omitempty"`
	AvailableUntil       *int64  `json:"available_until,omitempty"`
	IsVisible            bool    `json:"is_visible"`
	IsClosed             bool    `json:"is_closed"`
}

var (
	upsertKey         = "payment.product.upsert"
	upsertDescription = `
Creates or updates a product. Requires the 'payment.product.upsert'
permission in the target workspace.`
)

var Upsert = adapter.Method[UpsertRequest, struct{}]{
	Key:         upsertKey,
	Description: upsertDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(upsertKey)},
	Handler: func(ctx *adapter.Context, d UpsertRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.UpsertProduct(
			ctx.Context,
			padm.ProductUpsertParams{
				WorkspaceID:          d.WorkspaceID,
				ID:                   d.ID,
				GroupCode:            d.GroupCode,
				TitleKey:             d.TitleKey,
				DescriptionKey:       d.DescriptionKey,
				Target:               json.RawMessage(d.Target),
				ImageURL:             d.ImageURL,
				LinkURL:              d.LinkURL,
				SizeLabel:            d.SizeLabel,
				PeriodSeconds:        d.PeriodSeconds,
				TrialDurationSeconds: d.TrialDurationSeconds,
				QuantityMode:         d.QuantityMode,
				Position:             d.Position,
				GlobalLimit:          d.GlobalLimit,
				GlobalInterval:       d.GlobalInterval,
				GlobalIntervalCount:  d.GlobalIntervalCount,
				UserLimit:            d.UserLimit,
				UserInterval:         d.UserInterval,
				UserIntervalCount:    d.UserIntervalCount,
				AvailableFrom:        tsToTime(d.AvailableFrom),
				AvailableUntil:       tsToTime(d.AvailableUntil),
				IsVisible:            d.IsVisible,
				IsClosed:             d.IsClosed,
			})
	},
}

func tsToTime(ts *int64) *time.Time {
	if ts == nil {
		return nil
	}

	t := time.Unix(*ts, 0)

	return &t
}
