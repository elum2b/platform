package catalog

import (
	"encoding/json"
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type SaveProductRequest struct {
	WorkspaceID          string     `json:"workspace_id"              validate:"required,uuid"`
	ID                   string     `json:"id,omitempty"`
	GroupCode            *string    `json:"group_code,omitempty"`
	TitleKey             string     `json:"title_key"                 validate:"required,max=255"`
	DescriptionKey       *string    `json:"description_key,omitempty"`
	Target               string     `json:"target"                    validate:"required"`
	ImageURL             *string    `json:"image_url,omitempty"`
	LinkURL              *string    `json:"link_url,omitempty"`
	SizeLabel            *string    `json:"size_label,omitempty"`
	PeriodSeconds        *int64     `json:"period_seconds,omitempty"`
	TrialDurationSeconds *int64     `json:"trial_duration_seconds,omitempty"`
	QuantityMode         string     `json:"quantity_mode"             validate:"required"`
	Position             int32      `json:"position"`
	GlobalLimit          int32      `json:"global_limit"`
	GlobalInterval       string     `json:"global_interval"`
	GlobalIntervalCount  int32      `json:"global_interval_count"`
	UserLimit            int32      `json:"user_limit"`
	UserInterval         string     `json:"user_interval"`
	UserIntervalCount    int32      `json:"user_interval_count"`
	AvailableFrom        *time.Time `json:"available_from,omitempty"`
	AvailableUntil       *time.Time `json:"available_until,omitempty"`
	IsVisible            bool       `json:"is_visible"`
	IsClosed             bool       `json:"is_closed"`
}

var (
	saveProductKey         = "payment.catalog.save_product"
	saveProductDescription = `
Saves a product as part of catalog management. Requires the
'payment.catalog.save_product' permission in the target workspace.`
)

var SaveProduct = adapter.Method[SaveProductRequest, struct{}]{
	Key:         saveProductKey,
	Description: saveProductDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware:  []adapter.Middleware{adapter.WorkspaceAccess(saveProductKey)},
	Handler: func(ctx *adapter.Context, d SaveProductRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.SaveProduct(
			ctx.Context,
			padm.SaveProductParams{
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
				AvailableFrom:        d.AvailableFrom,
				AvailableUntil:       d.AvailableUntil,
				IsVisible:            d.IsVisible,
				IsClosed:             d.IsClosed,
			})
	},
}
