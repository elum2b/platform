package control

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/global"
	service "github.com/elum2b/services/control"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		global.Control = service.New(service.DatabaseParams{
			Host:     config.ControlPostgresHost,
			Port:     config.ControlPostgresPort,
			User:     config.ControlPostgresUser,
			Password: config.ControlPostgresPassword,
			Database: config.ControlPostgresDatabase,
			Options: service.Options{
				MaxConnections:      config.ControlMaxConnections,
				QueryTimeout:        config.ControlQueryTimeout,
				CacheL1Delay:        config.ControlCacheL1Delay,
				CacheL2Delay:        config.ControlCacheL2Delay,
				CacheEnabled:        config.ControlCacheEnabled,
				CacheSize:           config.ControlCacheSize,
				CacheTTLCheck:       config.ControlCacheTTLCheck,
				SecretEncryptionKey: []byte(config.ControlSecretEncryptionKey),
			},
		})

		return global.Control.Run(ctx)

	}

}
