package cpa

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/global"
	service "github.com/elum2b/services/cpa"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		global.CPA = service.New(service.DatabaseParams{
			Host:     config.CPAPostgresHost,
			Port:     config.CPAPostgresPort,
			User:     config.CPAPostgresUser,
			Password: config.CPAPostgresPassword,
			Database: config.CPAPostgresDatabase,
			Options: service.Options{
				MaxConnections: config.CPAMaxConnections,
				QueryTimeout:   config.CPAQueryTimeout,
				CacheL1Delay:   config.CPACacheL1Delay,
				CacheL2Delay:   config.CPACacheL2Delay,
				CacheEnabled:   config.CPACacheEnabled,
				CacheSize:      config.CPACacheSize,
				CacheTTLCheck:  config.CPACacheTTLCheck,
			},
		})

		if err := global.CPA.OnCallback(ctx, handler); err != nil {
			return err
		}

		return global.CPA.Run(ctx)

	}

}
