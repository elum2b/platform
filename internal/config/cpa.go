package config

import "github.com/elum-utils/env"

var (
	
	// CPAPostgresHost contains CPA PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	CPAPostgresHost = env.GetEnvString("CPA_POSTGRES_HOST", PostgresHost)

	// CPAPostgresPort contains CPA PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	CPAPostgresPort = env.GetEnvInt("CPA_POSTGRES_PORT", PostgresPort)

	// CPAPostgresUser contains CPA PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	CPAPostgresUser = env.GetEnvString("CPA_POSTGRES_USER", PostgresUser)

	// CPAPostgresPassword contains CPA PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	CPAPostgresPassword = env.GetEnvString("CPA_POSTGRES_PASSWORD", PostgresPassword)

	// CPAPostgresDatabase contains CPA PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	CPAPostgresDatabase = env.GetEnvString("CPA_POSTGRES_DATABASE", PostgresDatabase)

	// CPAMaxConnections contains CPA connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	CPAMaxConnections = env.GetEnvInt("CPA_MAX_CONNECTIONS", ServicesMaxConnections)

	// CPAQueryTimeout contains CPA query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	CPAQueryTimeout = env.GetEnvDuration("CPA_QUERY_TIMEOUT", ServicesQueryTimeout)

	// CPACacheL1Delay contains CPA in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	CPACacheL1Delay = env.GetEnvDuration("CPA_CACHE_L1_DELAY", ServicesCacheL1Delay)

	// CPACacheL2Delay contains CPA shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	CPACacheL2Delay = env.GetEnvDuration("CPA_CACHE_L2_DELAY", ServicesCacheL2Delay)

	// CPACacheEnabled contains CPA cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	CPACacheEnabled = env.GetEnvBool("CPA_CACHE_ENABLED", ServicesCacheEnabled)

	// CPACacheSize contains CPA in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	CPACacheSize = env.GetEnvInt("CPA_CACHE_SIZE", ServicesCacheSize)

	// CPACacheTTLCheck contains CPA cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	CPACacheTTLCheck = env.GetEnvDuration("CPA_CACHE_TTL_CHECK", ServicesCacheTTLCheck)

)
