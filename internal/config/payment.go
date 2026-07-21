package config

import "github.com/elum-utils/env"

var (

	// PaymentPostgresHost contains Payment PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	PaymentPostgresHost = env.GetEnvString("PAYMENT_POSTGRES_HOST", PostgresHost)

	// PaymentPostgresPort contains Payment PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	PaymentPostgresPort = env.GetEnvInt("PAYMENT_POSTGRES_PORT", PostgresPort)

	// PaymentPostgresUser contains Payment PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	PaymentPostgresUser = env.GetEnvString("PAYMENT_POSTGRES_USER", PostgresUser)

	// PaymentPostgresPassword contains Payment PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	PaymentPostgresPassword = env.GetEnvString("PAYMENT_POSTGRES_PASSWORD", PostgresPassword)

	// PaymentPostgresDatabase contains Payment PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	PaymentPostgresDatabase = env.GetEnvString("PAYMENT_POSTGRES_DATABASE", PostgresDatabase)

	// PaymentMaxConnections contains Payment connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	PaymentMaxConnections = env.GetEnvInt("PAYMENT_MAX_CONNECTIONS", ServicesMaxConnections)

	// PaymentQueryTimeout contains Payment query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	PaymentQueryTimeout = env.GetEnvDuration("PAYMENT_QUERY_TIMEOUT", ServicesQueryTimeout)

	// PaymentCacheL1Delay contains Payment in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	PaymentCacheL1Delay = env.GetEnvDuration("PAYMENT_CACHE_L1_DELAY", ServicesCacheL1Delay)

	// PaymentCacheL2Delay contains Payment shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	PaymentCacheL2Delay = env.GetEnvDuration("PAYMENT_CACHE_L2_DELAY", ServicesCacheL2Delay)

	// PaymentCacheEnabled contains Payment cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	PaymentCacheEnabled = env.GetEnvBool("PAYMENT_CACHE_ENABLED", ServicesCacheEnabled)

	// PaymentCacheSize contains Payment in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	PaymentCacheSize = env.GetEnvInt("PAYMENT_CACHE_SIZE", ServicesCacheSize)

	// PaymentCacheTTLCheck contains Payment cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	PaymentCacheTTLCheck = env.GetEnvDuration("PAYMENT_CACHE_TTL_CHECK", ServicesCacheTTLCheck)

	// PaymentPriceUpdateInterval contains price update interval.
	// Env: PAYMENT_PRICE_UPDATE_INTERVAL.
	PaymentPriceUpdateInterval = env.GetEnvDuration("PAYMENT_PRICE_UPDATE_INTERVAL", 0)

	// PaymentPriceUpdateBaseURL contains price provider base URL.
	// Env: PAYMENT_PRICE_UPDATE_BASE_URL.
	PaymentPriceUpdateBaseURL = env.GetEnvString("PAYMENT_PRICE_UPDATE_BASE_URL", "")

	// PaymentDisablePriceUpdater contains price updater state.
	// Env: PAYMENT_DISABLE_PRICE_UPDATER.
	PaymentDisablePriceUpdater = env.GetEnvBool("PAYMENT_DISABLE_PRICE_UPDATER", false)


	// PaymentOrderExpirationInterval contains order expiration interval.
	// Env: PAYMENT_ORDER_EXPIRATION_INTERVAL.
	PaymentOrderExpirationInterval = env.GetEnvDuration("PAYMENT_ORDER_EXPIRATION_INTERVAL", 0)

	// PaymentOrderExpirationAge contains maximum pending order age.
	// Env: PAYMENT_ORDER_EXPIRATION_AGE.
	PaymentOrderExpirationAge = env.GetEnvDuration("PAYMENT_ORDER_EXPIRATION_AGE", 0)

	// PaymentOrderExpirationBatch contains order expiration batch size.
	// Env: PAYMENT_ORDER_EXPIRATION_BATCH.
	PaymentOrderExpirationBatch = int32(env.GetEnvInt("PAYMENT_ORDER_EXPIRATION_BATCH", 0))

	// PaymentDisableOrderExpiration contains order expiration worker state.
	// Env: PAYMENT_DISABLE_ORDER_EXPIRATION.
	PaymentDisableOrderExpiration = env.GetEnvBool("PAYMENT_DISABLE_ORDER_EXPIRATION", false)

	// PaymentPlategaReconcileInterval contains Platega reconciliation interval.
	// Env: PAYMENT_PLATEGA_RECONCILE_INTERVAL.
	PaymentPlategaReconcileInterval = env.GetEnvDuration("PAYMENT_PLATEGA_RECONCILE_INTERVAL", 0)

	// PaymentPlategaReconcileMinAge contains minimum reconciliation age.
	// Env: PAYMENT_PLATEGA_RECONCILE_MIN_AGE.
	PaymentPlategaReconcileMinAge = env.GetEnvDuration("PAYMENT_PLATEGA_RECONCILE_MIN_AGE", 0)

	// PaymentPlategaReconcileMissingAfter contains missing payment threshold.
	// Env: PAYMENT_PLATEGA_RECONCILE_MISSING_AFTER.
	PaymentPlategaReconcileMissingAfter = env.GetEnvDuration("PAYMENT_PLATEGA_RECONCILE_MISSING_AFTER", 0)

	// PaymentPlategaReconcileBatch contains reconciliation batch size.
	// Env: PAYMENT_PLATEGA_RECONCILE_BATCH.
	PaymentPlategaReconcileBatch = int32(env.GetEnvInt("PAYMENT_PLATEGA_RECONCILE_BATCH", 0))

	// PaymentTONWalletSyncInterval contains TON wallet synchronization interval.
	// Env: PAYMENT_TON_WALLET_SYNC_INTERVAL.
	PaymentTONWalletSyncInterval = env.GetEnvDuration("PAYMENT_TON_WALLET_SYNC_INTERVAL", 0)

)
