package cpa

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"

	"github.com/elum2b/services/cpa"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		services.CPA = cpa.New(cpa.DatabaseParams{
			Host:     config.CPAPostgresHost,
			Port:     config.CPAPostgresPort,
			User:     config.CPAPostgresUser,
			Password: config.CPAPostgresPassword,
			Database: config.CPAPostgresDatabase,
			Options: cpa.Options{
				MaxConnections: config.CPAMaxConnections,
				QueryTimeout:   config.CPAQueryTimeout,
				CacheL1Delay:   config.CPACacheL1Delay,
				CacheL2Delay:   config.CPACacheL2Delay,
				CacheEnabled:   config.CPACacheEnabled,
				CacheSize:      config.CPACacheSize,
				CacheTTLCheck:  config.CPACacheTTLCheck,
			},
		})

		if err := services.CPA.OnCallback(ctx, handler); err != nil {
			return err
		}

		return services.CPA.Run(ctx)

	}

}
