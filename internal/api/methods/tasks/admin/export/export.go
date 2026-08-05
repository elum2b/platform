package exportapi

import (
	"time"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

type ManifestResponse struct {
	Manifest tadmin.ExportManifest `json:"manifest"`
}

var (
	manifestKey         = "tasks.export.manifest"
	manifestDescription = `
Returns the export manifest describing the sections available for export.
Requires the 'tasks.export.manifest' permission in the target workspace.`
)

// Manifest exposes the export manifest method.
var Manifest = adapter.Method[struct{}, ManifestResponse]{
	Key:         manifestKey,
	Description: manifestDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.GlobalAccess(manifestKey),
	},
	Handler: func(ctx *adapter.Context, _ struct{}) (ManifestResponse, error) {
		value, err := services.Tasks.Admin.ExportManifest(
			ctx.Context,
		)

		return ManifestResponse{Manifest: value}, err
	},
}

type Request struct {
	WorkspaceID string    `json:"workspace_id"  validate:"required,uuid"`
	Now         time.Time `json:"now,omitempty"`
}

type Response struct {
	Package tadmin.ExportPackage `json:"package"`
}

var (
	methodKey         = "tasks.export"
	methodDescription = `
Exports all tasks and their configuration from a workspace. Requires the
'tasks.export' permission in the target workspace.`
)

// Method exposes the tasks export method.
var Method = adapter.Method[Request, Response]{
	Key:         methodKey,
	Description: methodDescription,
	Transports:  adapter.WS | adapter.MCP,
	Middleware: []adapter.Middleware{
		adapter.WorkspaceAccess(methodKey),
	},
	Handler: func(ctx *adapter.Context, data Request) (Response, error) {
		value, err := services.Tasks.Admin.Export(
			ctx.Context,
			data.WorkspaceID,
			tadmin.ExportRequest{Now: data.Now},
		)

		return Response{Package: value}, err
	},
}
