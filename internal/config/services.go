package config

import "github.com/elum-utils/env"

var (
	// ServicesMaxConnections contains the default service connection limit.
	// Env: SERVICES_MAX_CONNECTIONS.
	ServicesMaxConnections = env.GetEnvInt("SERVICES_MAX_CONNECTIONS", 0)

	// ServicesQueryTimeout contains the default service query timeout.
	// Env: SERVICES_QUERY_TIMEOUT.
	ServicesQueryTimeout = env.GetEnvDuration("SERVICES_QUERY_TIMEOUT", 0)

	// ServicesCacheL1Delay contains the default in-memory cache lifetime.
	// Env: SERVICES_CACHE_L1_DELAY.
	ServicesCacheL1Delay = env.GetEnvDuration("SERVICES_CACHE_L1_DELAY", 0)

	// ServicesCacheL2Delay contains the default shared cache lifetime.
	// Env: SERVICES_CACHE_L2_DELAY.
	ServicesCacheL2Delay = env.GetEnvDuration("SERVICES_CACHE_L2_DELAY", 0)

	// ServicesCacheEnabled contains the default cache state.
	// Env: SERVICES_CACHE_ENABLED.
	ServicesCacheEnabled = env.GetEnvBool("SERVICES_CACHE_ENABLED", false)

	// ServicesCacheSize contains the default in-memory cache size.
	// Env: SERVICES_CACHE_SIZE.
	ServicesCacheSize = env.GetEnvInt("SERVICES_CACHE_SIZE", 0)

	// ServicesCacheTTLCheck contains the default cache expiry check interval.
	// Env: SERVICES_CACHE_TTL_CHECK.
	ServicesCacheTTLCheck = env.GetEnvDuration("SERVICES_CACHE_TTL_CHECK", 0)
)
