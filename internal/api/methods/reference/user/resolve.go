package user

import (
	refuser "github.com/elum2b/services/reference/service/user"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ResolveRequest struct {
	WorkspaceID string   `json:"workspace_id"     query:"workspace_id" validate:"required,uuid"`
	AppID       int64    `json:"app_id"           query:"app_id"       validate:"required,min=1"`
	PlatformID  int64    `json:"platform_id"      query:"platform_id"  validate:"required,min=1"`
	Keys        []string `json:"keys"             query:"keys"         validate:"required,min=1,max=1000"`
	Locale      string   `json:"locale,omitempty" query:"locale"`
}

type ResolveResponse struct {
	Result refuser.ResolveResult `json:"result"`
}

var (
	resolveKey         = "reference.user.resolve"
	resolveDescription = `
Resolves multiple reference items by key for the authenticated application
user. Returns found items and a list of missing keys.`
)

// Resolve exposes the reference user resolve method.
var Resolve = adapter.Method[ResolveRequest, ResolveResponse]{
	Key:         resolveKey,
	Description: resolveDescription,
	Transports:  adapter.HTTP,
	Handler: func(ctx *adapter.Context, data ResolveRequest) (ResolveResponse, error) {
		result, err := services.Reference.User.Resolve(
			ctx.Context,
			refuser.ResolveParams{
				WorkspaceID: data.WorkspaceID,
				Keys:        data.Keys,
				Locale:      data.Locale,
			},
		)

		return ResolveResponse{Result: result}, err
	},
}
