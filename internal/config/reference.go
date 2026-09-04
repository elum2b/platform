package config

import "github.com/elum-utils/env"

var (
	// ReferencePostgresHost contains Reference PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	ReferencePostgresHost = env.GetEnvString(
		"REFERENCE_POSTGRES_HOST",
		PostgresHost,
	)

	// ReferencePostgresPort contains Reference PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	ReferencePostgresPort = env.GetEnvInt(
		"REFERENCE_POSTGRES_PORT",
		PostgresPort,
	)

	// ReferencePostgresUser contains Reference PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	ReferencePostgresUser = env.GetEnvString(
		"REFERENCE_POSTGRES_USER",
		PostgresUser,
	)

	// ReferencePostgresPassword contains Reference PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	ReferencePostgresPassword = env.GetEnvString(
		"REFERENCE_POSTGRES_PASSWORD",
		PostgresPassword,
	)

	// ReferencePostgresDatabase contains Reference PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	ReferencePostgresDatabase = env.GetEnvString(
		"REFERENCE_POSTGRES_DATABASE",
		PostgresDatabase,
	)

	// ReferenceMaxConnections contains Reference connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	ReferenceMaxConnections = env.GetEnvInt(
		"REFERENCE_MAX_CONNECTIONS",
		ServicesMaxConnections,
	)

	// ReferenceQueryTimeout contains Reference query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	ReferenceQueryTimeout = env.GetEnvDuration(
		"REFERENCE_QUERY_TIMEOUT",
		ServicesQueryTimeout,
	)

	// ReferenceCacheL1Delay contains Reference in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	ReferenceCacheL1Delay = env.GetEnvDuration(
		"REFERENCE_CACHE_L1_DELAY",
		ServicesCacheL1Delay,
	)

	// ReferenceCacheL2Delay contains Reference shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	ReferenceCacheL2Delay = env.GetEnvDuration(
		"REFERENCE_CACHE_L2_DELAY",
		ServicesCacheL2Delay,
	)

	// ReferenceCacheEnabled contains Reference cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	ReferenceCacheEnabled = env.GetEnvBool(
		"REFERENCE_CACHE_ENABLED",
		ServicesCacheEnabled,
	)

	// ReferenceCacheSize contains Reference in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	ReferenceCacheSize = env.GetEnvInt(
		"REFERENCE_CACHE_SIZE",
		ServicesCacheSize,
	)

	// ReferenceCacheTTLCheck contains Reference cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	ReferenceCacheTTLCheck = env.GetEnvDuration(
		"REFERENCE_CACHE_TTL_CHECK",
		ServicesCacheTTLCheck,
	)

	// ReferenceStorageDirectory contains Reference local resource directory.
	ReferenceStorageDirectory = env.GetEnvString(
		"REFERENCE_STORAGE_DIRECTORY",
		"",
	)

	// ReferenceStorageEndpoint contains Reference S3/MinIO endpoint.
	// Fallback: **S3_ENDPOINT**.
	ReferenceStorageEndpoint = env.GetEnvString(
		"REFERENCE_STORAGE_ENDPOINT",
		S3Endpoint,
	)

	// ReferenceStorageBucket contains Reference S3/MinIO bucket.
	// When empty, local disk storage remains active.
	// Fallback: **S3_BUCKET**.
	ReferenceStorageBucket = env.GetEnvString(
		"REFERENCE_STORAGE_BUCKET",
		S3Bucket,
	)

	// ReferenceStorageAccessKey contains Reference S3/MinIO access key.
	// Fallback: **S3_ACCESS_KEY**.
	ReferenceStorageAccessKey = env.GetEnvString(
		"REFERENCE_STORAGE_ACCESS_KEY",
		S3AccessKey,
	)

	// ReferenceStorageSecretKey contains Reference S3/MinIO secret key.
	// Fallback: **S3_SECRET_KEY**.
	ReferenceStorageSecretKey = env.GetEnvString(
		"REFERENCE_STORAGE_SECRET_KEY",
		S3SecretKey,
	)

	// ReferenceStorageSessionToken contains Reference S3 session token.
	// Fallback: **S3_SESSION_TOKEN**.
	ReferenceStorageSessionToken = env.GetEnvString(
		"REFERENCE_STORAGE_SESSION_TOKEN",
		S3SessionToken,
	)

	// ReferenceStorageRegion contains Reference S3/MinIO region.
	// Fallback: **S3_REGION**.
	ReferenceStorageRegion = env.GetEnvString(
		"REFERENCE_STORAGE_REGION",
		S3Region,
	)

	// ReferenceStorageSecure contains Reference S3/MinIO HTTPS state.
	// Fallback: **S3_SECURE**.
	ReferenceStorageSecure = env.GetEnvBool(
		"REFERENCE_STORAGE_SECURE",
		S3Secure,
	)

	// ReferenceStorageUsePathStyle contains Reference S3 path-style state.
	// Fallback: **S3_USE_PATH_STYLE**.
	ReferenceStorageUsePathStyle = env.GetEnvBool(
		"REFERENCE_STORAGE_USE_PATH_STYLE",
		S3UsePathStyle,
	)
)
