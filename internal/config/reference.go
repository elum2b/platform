package config

import "github.com/elum-utils/env"

var (
	// ReferencePostgresHost contains Reference PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	ReferencePostgresHost = env.GetEnvString("REFERENCE_POSTGRES_HOST", PostgresHost)

	// ReferencePostgresPort contains Reference PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	ReferencePostgresPort = env.GetEnvInt("REFERENCE_POSTGRES_PORT", PostgresPort)

	// ReferencePostgresUser contains Reference PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	ReferencePostgresUser = env.GetEnvString("REFERENCE_POSTGRES_USER", PostgresUser)

	// ReferencePostgresPassword contains Reference PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	ReferencePostgresPassword = env.GetEnvString("REFERENCE_POSTGRES_PASSWORD", PostgresPassword)

	// ReferencePostgresDatabase contains Reference PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	ReferencePostgresDatabase = env.GetEnvString("REFERENCE_POSTGRES_DATABASE", PostgresDatabase)

	// ReferenceMaxConnections contains Reference connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	ReferenceMaxConnections = env.GetEnvInt("REFERENCE_MAX_CONNECTIONS", ServicesMaxConnections)

	// ReferenceQueryTimeout contains Reference query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	ReferenceQueryTimeout = env.GetEnvDuration("REFERENCE_QUERY_TIMEOUT", ServicesQueryTimeout)

	// ReferenceCacheL1Delay contains Reference in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	ReferenceCacheL1Delay = env.GetEnvDuration("REFERENCE_CACHE_L1_DELAY", ServicesCacheL1Delay)

	// ReferenceCacheL2Delay contains Reference shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	ReferenceCacheL2Delay = env.GetEnvDuration("REFERENCE_CACHE_L2_DELAY", ServicesCacheL2Delay)

	// ReferenceCacheEnabled contains Reference cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	ReferenceCacheEnabled = env.GetEnvBool("REFERENCE_CACHE_ENABLED", ServicesCacheEnabled)

	// ReferenceCacheSize contains Reference in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	ReferenceCacheSize = env.GetEnvInt("REFERENCE_CACHE_SIZE", ServicesCacheSize)

	// ReferenceCacheTTLCheck contains Reference cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	ReferenceCacheTTLCheck = env.GetEnvDuration("REFERENCE_CACHE_TTL_CHECK", ServicesCacheTTLCheck)
	
)
