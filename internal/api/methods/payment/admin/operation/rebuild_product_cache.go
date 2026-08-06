package operation

import (
	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type RebuildProductCacheRequest struct {
	WorkspaceID string `json:"workspace_id" validate:"required,uuid"`
}

var (
	rebuildProductCacheKey         = "payment.operation.rebuild_product_cache"
	rebuildProductCacheDescription = `
Rebuilds the product cache. Requires the
'payment.operation.rebuild_product_cache' permission in the target workspace.`
)

var RebuildProductCache = adapter.Method[RebuildProductCacheRequest, struct{}]{
	Key:         rebuildProductCacheKey,
	Description: rebuildProductCacheDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(rebuildProductCacheKey),
	},
	Handler: func(ctx *adapter.Context, d RebuildProductCacheRequest) (struct{}, error) {
		return struct{}{}, services.Payment.Admin.RebuildProductCache(
			ctx.Context, d.WorkspaceID)
	},
}
