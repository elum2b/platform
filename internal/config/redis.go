package config

import "github.com/elum-utils/env"

var (

	// RedisHost contains the shared Redis host reserved for Redis integration.
	// Env: REDIS_HOST.
	RedisHost = env.GetEnvString(
		"REDIS_HOST",
		"localhost",
	)

	// RedisPort contains the shared Redis port reserved for Redis integration.
	// Env: REDIS_PORT.
	RedisPort = env.GetEnvInt(
		"REDIS_PORT",
		6379,
	)

	// RedisUser contains the shared Redis user reserved for Redis integration.
	// Env: REDIS_USER.
	RedisUser = env.GetEnvString(
		"REDIS_USER",
		"",
	)

	// RedisPassword contains the shared Redis password reserved for Redis integration.
	// Env: REDIS_PASSWORD.
	RedisPassword = env.GetEnvString(
		"REDIS_PASSWORD",
		"",
	)

	// RedisDatabase contains the shared Redis database reserved for Redis integration.
	// Env: REDIS_DATABASE.
	RedisDatabase = env.GetEnvInt(
		"REDIS_DATABASE",
		0,
	)

	// RedisSecure contains the shared Redis TLS state reserved for Redis integration.
	// Env: REDIS_SECURE.
	RedisSecure = env.GetEnvBool(
		"REDIS_SECURE",
		false,
	)
)
