package config

import (
	"time"

	"github.com/elum-utils/env"
)

var (
	// Host contains the server host.
	// Env: HOST.
	Host = env.GetEnvString(
		"HOST",
		"0.0.0.0",
	)

	// Port contains the server port.
	// Env: PORT.
	Port = env.GetEnvInt(
		"PORT",
		18300,
	)

	// SystemVersionRepository contains the OCI repository with releases.
	// Env: SYSTEM_VERSION_REPOSITORY.
	SystemVersionRepository = env.GetEnvString(
		"SYSTEM_VERSION_REPOSITORY",
		"elum2b/platform",
	)

	// SystemVersionCacheDuration contains the latest-version cache lifetime.
	// Env: SYSTEM_VERSION_CACHE_DURATION.
	SystemVersionCacheDuration = env.GetEnvDuration(
		"SYSTEM_VERSION_CACHE_DURATION",
		10*time.Minute,
	)
)
