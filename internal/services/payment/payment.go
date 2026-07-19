package payment

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/global"
	service "github.com/elum2b/services/payment"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		global.Payment = service.New(service.DatabaseParams{
			Host:     config.PaymentPostgresHost,
			Port:     config.PaymentPostgresPort,
			User:     config.PaymentPostgresUser,
			Password: config.PaymentPostgresPassword,
			Database: config.PaymentPostgresDatabase,
			Options: service.Options{
				MaxConnections: config.PaymentMaxConnections,
				QueryTimeout:   config.PaymentQueryTimeout,
				CacheL1Delay:   config.PaymentCacheL1Delay,
				CacheL2Delay:   config.PaymentCacheL2Delay,
				CacheEnabled:   config.PaymentCacheEnabled,
				CacheSize:      config.PaymentCacheSize,
				CacheTTLCheck:  config.PaymentCacheTTLCheck,

				PriceUpdateInterval: config.PaymentPriceUpdateInterval,
				PriceUpdateBaseURL:  config.PaymentPriceUpdateBaseURL,
				DisablePriceUpdater: config.PaymentDisablePriceUpdater,

				OrderExpirationInterval: config.PaymentOrderExpirationInterval,
				OrderExpirationAge:      config.PaymentOrderExpirationAge,
				OrderExpirationBatch:    config.PaymentOrderExpirationBatch,
				DisableOrderExpiration:  config.PaymentDisableOrderExpiration,

				PlategaReconcileInterval:     config.PaymentPlategaReconcileInterval,
				PlategaReconcileMinAge:       config.PaymentPlategaReconcileMinAge,
				PlategaReconcileMissingAfter: config.PaymentPlategaReconcileMissingAfter,
				PlategaReconcileBatch:        config.PaymentPlategaReconcileBatch,

				TONWalletSyncInterval: config.PaymentTONWalletSyncInterval,
			},
		})

		if err := global.Payment.OnCallback(ctx, handler); err != nil {
			return err
		}

		return global.Payment.Run(ctx)

	}

}
