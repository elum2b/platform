package config

import "github.com/elum-utils/env"

var (
	// TasksPostgresHost contains Tasks PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	TasksPostgresHost = env.GetEnvString("TASKS_POSTGRES_HOST", PostgresHost)
	// TasksPostgresPort contains Tasks PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	TasksPostgresPort = env.GetEnvInt("TASKS_POSTGRES_PORT", PostgresPort)
	// TasksPostgresUser contains Tasks PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	TasksPostgresUser = env.GetEnvString("TASKS_POSTGRES_USER", PostgresUser)
	// TasksPostgresPassword contains Tasks PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	TasksPostgresPassword = env.GetEnvString("TASKS_POSTGRES_PASSWORD", PostgresPassword)
	// TasksPostgresDatabase contains Tasks PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	TasksPostgresDatabase = env.GetEnvString("TASKS_POSTGRES_DATABASE", PostgresDatabase)

	// TasksMaxConnections contains Tasks connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	TasksMaxConnections = env.GetEnvInt("TASKS_MAX_CONNECTIONS", ServicesMaxConnections)
	// TasksQueryTimeout contains Tasks query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	TasksQueryTimeout = env.GetEnvDuration("TASKS_QUERY_TIMEOUT", ServicesQueryTimeout)
	// TasksCacheL1Delay contains Tasks in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	TasksCacheL1Delay = env.GetEnvDuration("TASKS_CACHE_L1_DELAY", ServicesCacheL1Delay)
	// TasksCacheL2Delay contains Tasks shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	TasksCacheL2Delay = env.GetEnvDuration("TASKS_CACHE_L2_DELAY", ServicesCacheL2Delay)
	// TasksCacheEnabled contains Tasks cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	TasksCacheEnabled = env.GetEnvBool("TASKS_CACHE_ENABLED", ServicesCacheEnabled)
	// TasksCacheSize contains Tasks in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	TasksCacheSize = env.GetEnvInt("TASKS_CACHE_SIZE", ServicesCacheSize)
	// TasksCacheTTLCheck contains Tasks cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	TasksCacheTTLCheck = env.GetEnvDuration("TASKS_CACHE_TTL_CHECK", ServicesCacheTTLCheck)

	// TasksPartnerStartLeaseDuration contains partner task start lease duration.
	// Env: TASKS_PARTNER_START_LEASE_DURATION.
	TasksPartnerStartLeaseDuration = env.GetEnvDuration("TASKS_PARTNER_START_LEASE_DURATION", 0)

	// TasksRuntimeEnabled contains Lua runtime state.
	// Env: TASKS_RUNTIME_ENABLED.
	TasksRuntimeEnabled = env.GetEnvBool("TASKS_RUNTIME_ENABLED", false)
	// TasksRuntimeScriptCacheTTL contains Lua script cache lifetime.
	// Env: TASKS_RUNTIME_SCRIPT_CACHE_TTL.
	TasksRuntimeScriptCacheTTL = env.GetEnvDuration("TASKS_RUNTIME_SCRIPT_CACHE_TTL", 0)
	// TasksRuntimeTimeout contains Lua execution timeout.
	// Env: TASKS_RUNTIME_TIMEOUT.
	TasksRuntimeTimeout = env.GetEnvDuration("TASKS_RUNTIME_TIMEOUT", 0)
	// TasksRuntimeMaxMemory contains Lua runtime memory limit in bytes.
	// Env: TASKS_RUNTIME_MAX_MEMORY.
	TasksRuntimeMaxMemory = env.GetEnvInt("TASKS_RUNTIME_MAX_MEMORY", 0)
	// TasksRuntimeMaxHTTPRequests contains Lua HTTP request limit.
	// Env: TASKS_RUNTIME_MAX_HTTP_REQUESTS.
	TasksRuntimeMaxHTTPRequests = env.GetEnvInt("TASKS_RUNTIME_MAX_HTTP_REQUESTS", 0)
	// TasksRuntimeMaxResponseBytes contains Lua HTTP response limit in bytes.
	// Env: TASKS_RUNTIME_MAX_RESPONSE_BYTES.
	TasksRuntimeMaxResponseBytes = int64(env.GetEnvInt("TASKS_RUNTIME_MAX_RESPONSE_BYTES", 0))
	// TasksRuntimeJSONBoundary contains Lua JSON boundary state.
	// Env: TASKS_RUNTIME_JSON_BOUNDARY.
	TasksRuntimeJSONBoundary = env.GetEnvBool("TASKS_RUNTIME_JSON_BOUNDARY", false)
	// TasksRuntimeStatePoolSize contains Lua state pool size.
	// Env: TASKS_RUNTIME_STATE_POOL_SIZE.
	TasksRuntimeStatePoolSize = env.GetEnvInt("TASKS_RUNTIME_STATE_POOL_SIZE", 0)
)
