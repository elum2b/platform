package reference

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/global"
	service "github.com/elum2b/services/reference"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		global.Reference = service.New(service.DatabaseParams{
			Host:     config.ReferencePostgresHost,
			Port:     config.ReferencePostgresPort,
			User:     config.ReferencePostgresUser,
			Password: config.ReferencePostgresPassword,
			Database: config.ReferencePostgresDatabase,
			Options: service.Options{
				MaxConnections: config.ReferenceMaxConnections,
				QueryTimeout:   config.ReferenceQueryTimeout,
				CacheL1Delay:   config.ReferenceCacheL1Delay,
				CacheL2Delay:   config.ReferenceCacheL2Delay,
				CacheEnabled:   config.ReferenceCacheEnabled,
				CacheSize:      config.ReferenceCacheSize,
				CacheTTLCheck:  config.ReferenceCacheTTLCheck,
			},
		})

		return global.Reference.Run(ctx)

	}

}
