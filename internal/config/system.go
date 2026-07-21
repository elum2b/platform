package config

import "github.com/elum-utils/env"

var (
	
	// Host contains the server host.
	// Env: HOST.
	Host = env.GetEnvString("HOST", "0.0.0.0")

	// Port contains the server port.
	// Env: PORT.
	Port = env.GetEnvInt("PORT", 18300)

)
