package config

import "github.com/elum-utils/env"

var (
	// CalendarPostgresHost contains Calendar PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	CalendarPostgresHost = env.GetEnvString("CALENDAR_POSTGRES_HOST", PostgresHost)
	// CalendarPostgresPort contains Calendar PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	CalendarPostgresPort = env.GetEnvInt("CALENDAR_POSTGRES_PORT", PostgresPort)
	// CalendarPostgresUser contains Calendar PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	CalendarPostgresUser = env.GetEnvString("CALENDAR_POSTGRES_USER", PostgresUser)
	// CalendarPostgresPassword contains Calendar PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	CalendarPostgresPassword = env.GetEnvString("CALENDAR_POSTGRES_PASSWORD", PostgresPassword)
	// CalendarPostgresDatabase contains Calendar PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	CalendarPostgresDatabase = env.GetEnvString("CALENDAR_POSTGRES_DATABASE", PostgresDatabase)

	// CalendarMaxConnections contains Calendar connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	CalendarMaxConnections = env.GetEnvInt("CALENDAR_MAX_CONNECTIONS", ServicesMaxConnections)
	// CalendarQueryTimeout contains Calendar query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	CalendarQueryTimeout = env.GetEnvDuration("CALENDAR_QUERY_TIMEOUT", ServicesQueryTimeout)
	// CalendarCacheL1Delay contains Calendar in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	CalendarCacheL1Delay = env.GetEnvDuration("CALENDAR_CACHE_L1_DELAY", ServicesCacheL1Delay)
	// CalendarCacheL2Delay contains Calendar shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	CalendarCacheL2Delay = env.GetEnvDuration("CALENDAR_CACHE_L2_DELAY", ServicesCacheL2Delay)
	// CalendarCacheEnabled contains Calendar cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	CalendarCacheEnabled = env.GetEnvBool("CALENDAR_CACHE_ENABLED", ServicesCacheEnabled)
	// CalendarCacheSize contains Calendar in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	CalendarCacheSize = env.GetEnvInt("CALENDAR_CACHE_SIZE", ServicesCacheSize)
	// CalendarCacheTTLCheck contains Calendar cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	CalendarCacheTTLCheck = env.GetEnvDuration("CALENDAR_CACHE_TTL_CHECK", ServicesCacheTTLCheck)
)
