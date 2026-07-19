package config

import "github.com/elum-utils/env"

var (
	// ControlPostgresHost contains Control PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	ControlPostgresHost = env.GetEnvString("CONTROL_POSTGRES_HOST", PostgresHost)
	// ControlPostgresPort contains Control PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	ControlPostgresPort = env.GetEnvInt("CONTROL_POSTGRES_PORT", PostgresPort)
	// ControlPostgresUser contains Control PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	ControlPostgresUser = env.GetEnvString("CONTROL_POSTGRES_USER", PostgresUser)
	// ControlPostgresPassword contains Control PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	ControlPostgresPassword = env.GetEnvString("CONTROL_POSTGRES_PASSWORD", PostgresPassword)
	// ControlPostgresDatabase contains Control PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	ControlPostgresDatabase = env.GetEnvString("CONTROL_POSTGRES_DATABASE", PostgresDatabase)

	// ControlMaxConnections contains Control connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	ControlMaxConnections = env.GetEnvInt("CONTROL_MAX_CONNECTIONS", ServicesMaxConnections)
	// ControlQueryTimeout contains Control query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	ControlQueryTimeout = env.GetEnvDuration("CONTROL_QUERY_TIMEOUT", ServicesQueryTimeout)
	// ControlCacheL1Delay contains Control in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	ControlCacheL1Delay = env.GetEnvDuration("CONTROL_CACHE_L1_DELAY", ServicesCacheL1Delay)
	// ControlCacheL2Delay contains Control shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	ControlCacheL2Delay = env.GetEnvDuration("CONTROL_CACHE_L2_DELAY", ServicesCacheL2Delay)
	// ControlCacheEnabled contains Control cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	ControlCacheEnabled = env.GetEnvBool("CONTROL_CACHE_ENABLED", ServicesCacheEnabled)
	// ControlCacheSize contains Control in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	ControlCacheSize = env.GetEnvInt("CONTROL_CACHE_SIZE", ServicesCacheSize)
	// ControlCacheTTLCheck contains Control cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	ControlCacheTTLCheck = env.GetEnvDuration("CONTROL_CACHE_TTL_CHECK", ServicesCacheTTLCheck)

	// ControlSecretEncryptionKey contains the 32-byte Control encryption key.
	// Env: CONTROL_SECRET_ENCRYPTION_KEY.
	ControlSecretEncryptionKey = env.GetEnvString("CONTROL_SECRET_ENCRYPTION_KEY", "")
)
