package reference

import (
	"context"

	"github.com/elum2b/services/reference"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
)

func Service() func(context.Context) error {
	return func(ctx context.Context) error {
		return services.Reference.Run(ctx, reference.DatabaseParams{
			Host:     config.ReferencePostgresHost,
			Port:     config.ReferencePostgresPort,
			User:     config.ReferencePostgresUser,
			Password: config.ReferencePostgresPassword,
			Database: config.ReferencePostgresDatabase,
			SSLMode:  "disable",
			Options: reference.Options{
				MaxConnections: config.ReferenceMaxConnections,
				QueryTimeout:   config.ReferenceQueryTimeout,
				CacheL1Delay:   config.ReferenceCacheL1Delay,
				CacheL2Delay:   config.ReferenceCacheL2Delay,
				CacheEnabled:   config.ReferenceCacheEnabled,
				CacheSize:      config.ReferenceCacheSize,
				CacheTTLCheck:  config.ReferenceCacheTTLCheck,
			},
		})
	}
}
