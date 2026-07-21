package config

import "github.com/elum-utils/env"

var (
	
	// PostgresHost contains the default PostgreSQL host.
	// Env: POSTGRES_HOST.
	PostgresHost = env.GetEnvString("POSTGRES_HOST", "localhost")

	// PostgresPort contains the default PostgreSQL port.
	// Env: POSTGRES_PORT.
	PostgresPort = env.GetEnvInt("POSTGRES_PORT", 5432)

	// PostgresUser contains the default PostgreSQL user.
	// Env: POSTGRES_USER.
	PostgresUser = env.GetEnvString("POSTGRES_USER", "")

	// PostgresPassword contains the default PostgreSQL password.
	// Env: POSTGRES_PASSWORD.
	PostgresPassword = env.GetEnvString("POSTGRES_PASSWORD", "")

	// PostgresDatabase contains the default PostgreSQL database name.
	// Env: POSTGRES_DATABASE.
	PostgresDatabase = env.GetEnvString("POSTGRES_DATABASE", "")

)
