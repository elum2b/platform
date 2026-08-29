package exportapi

import (
	"time"

	tadmin "github.com/elum2b/services/tasks/service/admin"

	"github.com/elum2b/platform/internal/services"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	"github.com/elum2b/platform/internal/utils/archive"
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
	WorkspaceID    string    `json:"workspace_id"               validate:"required,uuid"`
	Sections       []string  `json:"sections,omitempty"`
	IncludeSecrets bool      `json:"include_secrets,omitempty"`
	Now            time.Time `json:"now,omitempty"`
}

type Response struct {
	Job tadmin.ArchiveJob `json:"job"`
}

var (
	methodKey         = "tasks.export"
	methodDescription = `
Queues an archive export of tasks and their configuration from a workspace. Requires the
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
		value, err := services.Tasks.Admin.QueueArchiveExport(
			ctx.Context,
			tadmin.QueueArchiveExportParams{
				WorkspaceID: data.WorkspaceID,
				FileName:    archive.FileName("tasks"),
				ExportRequest: tadmin.ExportRequest{
					Sections:       data.Sections,
					IncludeSecrets: data.IncludeSecrets,
					Now:            data.Now,
				},
			},
		)

		return Response{Job: value}, err
	},
}
