package promo

import (
	"context"

	"github.com/elum2b/services/promo"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
)

func Service() func(context.Context) error {
	return func(ctx context.Context) error {
		if err := services.Promo.OnCallback(ctx, handler); err != nil {
			return err
		}

		return services.Promo.Run(ctx, promo.DatabaseParams{
			Host:     config.PromoPostgresHost,
			Port:     config.PromoPostgresPort,
			User:     config.PromoPostgresUser,
			Password: config.PromoPostgresPassword,
			Database: config.PromoPostgresDatabase,
			Options: promo.Options{
				MaxConnections: config.PromoMaxConnections,
				QueryTimeout:   config.PromoQueryTimeout,
				CacheL1Delay:   config.PromoCacheL1Delay,
				CacheL2Delay:   config.PromoCacheL2Delay,
				CacheEnabled:   config.PromoCacheEnabled,
				CacheSize:      config.PromoCacheSize,
				CacheTTLCheck:  config.PromoCacheTTLCheck,
			},
		})
	}
}
