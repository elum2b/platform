package system

import (
	"net/http"
	"time"

	"github.com/elum2b/platform/internal/config"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	versionutils "github.com/elum2b/platform/internal/utils/version"
)

var versionClient = versionutils.NewClient(
	&http.Client{Timeout: 10 * time.Second},
	"https://ghcr.io",
	config.SystemVersionRepository,
	config.SystemVersionCacheDuration,
)

type VersionResponse struct {
	CurrentVersion string `json:"current_version"`
	LatestVersion  string `json:"latest_version"`
}

var (
	versionKey         = "system.version"
	versionDescription = `
Returns the installed platform version and the latest published platform version.
Requires an authenticated platform session or MCP token.`
)

// Version returns the installed and latest public platform versions.
var Version = adapter.Method[struct{}, VersionResponse]{
	Key:         versionKey,
	Description: versionDescription,
	Transports:  adapter.WS | adapter.MCP,
	Handler: func(
		ctx *adapter.Context,
		_ struct{},
	) (VersionResponse, error) {
		latest, err := versionClient.Latest(ctx.Context)
		if err != nil {
			return VersionResponse{}, err
		}

		return VersionResponse{
			CurrentVersion: versionutils.Current(),
			LatestVersion:  latest,
		}, nil
	},
}
