package product

import (
	"encoding/json"
	"time"

	padm "github.com/elum2b/services/payment/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type UpsertRequest struct {
	WorkspaceID          string  `json:"workspace_id"                     query:"workspace_id"           validate:"required,uuid"`
	ID                   string  `json:"id,omitempty"                     query:"id"`
	GroupCode            *string `json:"group_code,omitempty"             query:"group_code"`
	TitleKey             string  `json:"title_key"                        query:"title_key"              validate:"required,max=255"`
	DescriptionKey       *string `json:"description_key,omitempty"        query:"description_key"`
	Target               string  `json:"target"                           query:"target"                 validate:"required"`
	ImageURL             *string `json:"image_url,omitempty"              query:"image_url"`
	LinkURL              *string `json:"link_url,omitempty"               query:"link_url"`
	SizeLabel            *string `json:"size_label,omitempty"             query:"size_label"`
	PeriodSeconds        *int64  `json:"period_seconds,omitempty"         query:"period_seconds"`
	TrialDurationSeconds *int64  `json:"trial_duration_seconds,omitempty" query:"trial_duration_seconds"`
	QuantityMode         string  `json:"quantity_mode"                    query:"quantity_mode"          validate:"required"`
	Position             int32   `json:"position"                         query:"position"`
	GlobalLimit          int32   `json:"global_limit"                     query:"global_limit"`
	GlobalInterval       string  `json:"global_interval"                  query:"global_interval"`
	GlobalIntervalCount  int32   `json:"global_interval_count"            query:"global_interval_count"`
	UserLimit            int32   `json:"user_limit"                       query:"user_limit"`
	UserInterval         string  `json:"user_interval"                    query:"user_interval"`
	UserIntervalCount    int32   `json:"user_interval_count"              query:"user_interval_count"`
	AvailableFrom        *int64  `json:"available_from,omitempty"         query:"available_from"`
	AvailableUntil       *int64  `json:"available_until,omitempty"        query:"available_until"`
	IsVisible            bool    `json:"is_visible"                       query:"is_visible"`
	IsClosed             bool    `json:"is_closed"                        query:"is_closed"`
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
