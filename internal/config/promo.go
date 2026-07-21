package config

import "github.com/elum-utils/env"

var (

	// PromoPostgresHost contains Promo PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	PromoPostgresHost = env.GetEnvString("PROMO_POSTGRES_HOST", PostgresHost)

	// PromoPostgresPort contains Promo PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	PromoPostgresPort = env.GetEnvInt("PROMO_POSTGRES_PORT", PostgresPort)

	// PromoPostgresUser contains Promo PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	PromoPostgresUser = env.GetEnvString("PROMO_POSTGRES_USER", PostgresUser)

	// PromoPostgresPassword contains Promo PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	PromoPostgresPassword = env.GetEnvString("PROMO_POSTGRES_PASSWORD", PostgresPassword)

	// PromoPostgresDatabase contains Promo PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	PromoPostgresDatabase = env.GetEnvString("PROMO_POSTGRES_DATABASE", PostgresDatabase)

	// PromoMaxConnections contains Promo connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	PromoMaxConnections = env.GetEnvInt("PROMO_MAX_CONNECTIONS", ServicesMaxConnections)

	// PromoQueryTimeout contains Promo query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	PromoQueryTimeout = env.GetEnvDuration("PROMO_QUERY_TIMEOUT", ServicesQueryTimeout)

	// PromoCacheL1Delay contains Promo in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	PromoCacheL1Delay = env.GetEnvDuration("PROMO_CACHE_L1_DELAY", ServicesCacheL1Delay)

	// PromoCacheL2Delay contains Promo shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	PromoCacheL2Delay = env.GetEnvDuration("PROMO_CACHE_L2_DELAY", ServicesCacheL2Delay)

	// PromoCacheEnabled contains Promo cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	PromoCacheEnabled = env.GetEnvBool("PROMO_CACHE_ENABLED", ServicesCacheEnabled)

	// PromoCacheSize contains Promo in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	PromoCacheSize = env.GetEnvInt("PROMO_CACHE_SIZE", ServicesCacheSize)

	// PromoCacheTTLCheck contains Promo cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	PromoCacheTTLCheck = env.GetEnvDuration("PROMO_CACHE_TTL_CHECK", ServicesCacheTTLCheck)
	
)
