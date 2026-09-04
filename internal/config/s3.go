package config

import "github.com/elum-utils/env"

var (

	// S3Endpoint contains the default S3/MinIO endpoint.
	// Env: S3_ENDPOINT.
	S3Endpoint = env.GetEnvString(
		"S3_ENDPOINT",
		"",
	)

	// S3Bucket contains the default S3/MinIO bucket.
	// Env: S3_BUCKET.
	S3Bucket = env.GetEnvString(
		"S3_BUCKET",
		"",
	)

	// S3AccessKey contains the default S3/MinIO access key.
	// Env: S3_ACCESS_KEY.
	S3AccessKey = env.GetEnvString(
		"S3_ACCESS_KEY",
		"",
	)

	// S3SecretKey contains the default S3/MinIO secret key.
	// Env: S3_SECRET_KEY.
	S3SecretKey = env.GetEnvString(
		"S3_SECRET_KEY",
		"",
	)

	// S3SessionToken contains the default S3 session token.
	// Env: S3_SESSION_TOKEN.
	S3SessionToken = env.GetEnvString(
		"S3_SESSION_TOKEN",
		"",
	)

	// S3Region contains the default S3/MinIO region.
	// Env: S3_REGION.
	S3Region = env.GetEnvString(
		"S3_REGION",
		"",
	)

	// S3Secure contains the default S3/MinIO HTTPS state.
	// Env: S3_SECURE.
	S3Secure = env.GetEnvBool(
		"S3_SECURE",
		true,
	)

	// S3UsePathStyle contains the default S3 path-style state.
	// Env: S3_USE_PATH_STYLE.
	S3UsePathStyle = env.GetEnvBool(
		"S3_USE_PATH_STYLE",
		false,
	)
)
