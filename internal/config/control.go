package config

import (
	"time"

	"github.com/elum-utils/env"
)

var (
	// ControlPostgresHost contains Control PostgreSQL host.
	// Fallback: **POSTGRES_HOST**.
	ControlPostgresHost = env.GetEnvString(
		"CONTROL_POSTGRES_HOST",
		PostgresHost,
	)

	// ControlPostgresPort contains Control PostgreSQL port.
	// Fallback: **POSTGRES_PORT**.
	ControlPostgresPort = env.GetEnvInt(
		"CONTROL_POSTGRES_PORT",
		PostgresPort,
	)

	// ControlPostgresUser contains Control PostgreSQL user.
	// Fallback: **POSTGRES_USER**.
	ControlPostgresUser = env.GetEnvString(
		"CONTROL_POSTGRES_USER",
		PostgresUser,
	)

	// ControlPostgresPassword contains Control PostgreSQL password.
	// Fallback: **POSTGRES_PASSWORD**.
	ControlPostgresPassword = env.GetEnvString(
		"CONTROL_POSTGRES_PASSWORD",
		PostgresPassword,
	)

	// ControlPostgresDatabase contains Control PostgreSQL database name.
	// Fallback: **POSTGRES_DATABASE**.
	ControlPostgresDatabase = env.GetEnvString(
		"CONTROL_POSTGRES_DATABASE",
		PostgresDatabase,
	)

	// ControlMaxConnections contains Control connection limit.
	// Fallback: **SERVICES_MAX_CONNECTIONS**.
	ControlMaxConnections = env.GetEnvInt(
		"CONTROL_MAX_CONNECTIONS",
		ServicesMaxConnections,
	)

	// ControlQueryTimeout contains Control query timeout.
	// Fallback: **SERVICES_QUERY_TIMEOUT**.
	ControlQueryTimeout = env.GetEnvDuration(
		"CONTROL_QUERY_TIMEOUT",
		ServicesQueryTimeout,
	)

	// ControlCacheL1Delay contains Control in-memory cache lifetime.
	// Fallback: **SERVICES_CACHE_L1_DELAY**.
	ControlCacheL1Delay = env.GetEnvDuration(
		"CONTROL_CACHE_L1_DELAY",
		ServicesCacheL1Delay,
	)

	// ControlCacheL2Delay contains Control shared cache lifetime.
	// Fallback: **SERVICES_CACHE_L2_DELAY**.
	ControlCacheL2Delay = env.GetEnvDuration(
		"CONTROL_CACHE_L2_DELAY",
		ServicesCacheL2Delay,
	)

	// ControlCacheEnabled contains Control cache state.
	// Fallback: **SERVICES_CACHE_ENABLED**.
	ControlCacheEnabled = env.GetEnvBool(
		"CONTROL_CACHE_ENABLED",
		ServicesCacheEnabled,
	)

	// ControlCacheSize contains Control in-memory cache size.
	// Fallback: **SERVICES_CACHE_SIZE**.
	ControlCacheSize = env.GetEnvInt(
		"CONTROL_CACHE_SIZE",
		ServicesCacheSize,
	)

	// ControlCacheTTLCheck contains Control cache expiry check interval.
	// Fallback: **SERVICES_CACHE_TTL_CHECK**.
	ControlCacheTTLCheck = env.GetEnvDuration(
		"CONTROL_CACHE_TTL_CHECK",
		ServicesCacheTTLCheck,
	)

	// ControlSecretEncryptionKey contains the 32-byte Control encryption key.
	// Env: CONTROL_SECRET_ENCRYPTION_KEY.
	ControlSecretEncryptionKey = env.GetEnvString(
		"CONTROL_SECRET_ENCRYPTION_KEY",
		"",
	)

	// ControlAuthSessionDuration contains the Control authentication session lifetime.
	// Env: CONTROL_AUTH_SESSION_DURATION.
	ControlAuthSessionDuration = env.GetEnvDuration(
		"CONTROL_AUTH_SESSION_DURATION",
		48*time.Hour,
	)

	// ControlAuthCookieName contains the Control authentication cookie name.
	// Env: CONTROL_AUTH_COOKIE_NAME.
	ControlAuthCookieName = env.GetEnvString(
		"CONTROL_AUTH_COOKIE_NAME",
		"elum_session",
	)

	// ControlAuthCookieDomain contains the optional Control authentication cookie domain.
	// Env: CONTROL_AUTH_COOKIE_DOMAIN.
	ControlAuthCookieDomain = env.GetEnvString(
		"CONTROL_AUTH_COOKIE_DOMAIN",
		"",
	)

	// ControlAuthCookieSameSite contains the Control authentication cookie SameSite policy.
	// Env: CONTROL_AUTH_COOKIE_SAME_SITE.
	ControlAuthCookieSameSite = env.GetEnvString(
		"CONTROL_AUTH_COOKIE_SAME_SITE",
		"Strict",
	)

	// ControlAuthTwoFactorCookieName contains the Control two-factor challenge cookie name.
	// Env: CONTROL_AUTH_TWO_FACTOR_COOKIE_NAME.
	ControlAuthTwoFactorCookieName = env.GetEnvString(
		"CONTROL_AUTH_TWO_FACTOR_COOKIE_NAME",
		"elum_two_factor",
	)

	// ControlAuthTwoFactorDuration contains the Control two-factor challenge cookie lifetime.
	// Env: CONTROL_AUTH_TWO_FACTOR_DURATION.
	ControlAuthTwoFactorDuration = env.GetEnvDuration(
		"CONTROL_AUTH_TWO_FACTOR_DURATION",
		15*time.Minute,
	)

	// ControlAuthOAuthTimeout contains the Control OAuth provider request timeout.
	// Env: CONTROL_AUTH_OAUTH_TIMEOUT.
	ControlAuthOAuthTimeout = env.GetEnvDuration(
		"CONTROL_AUTH_OAUTH_TIMEOUT",
		15*time.Second,
	)

	// ControlAuthVKClientID contains the VK ID OAuth client ID.
	// Env: CONTROL_AUTH_VK_CLIENT_ID.
	ControlAuthVKClientID = env.GetEnvString(
		"CONTROL_AUTH_VK_CLIENT_ID",
		"",
	)

	// ControlAuthVKClientSecret contains the VK ID OAuth client secret.
	// Env: CONTROL_AUTH_VK_CLIENT_SECRET.
	ControlAuthVKClientSecret = env.GetEnvString(
		"CONTROL_AUTH_VK_CLIENT_SECRET",
		"",
	)

	// ControlAuthVKTokenURL contains the optional VK ID OAuth token endpoint.
	// Env: CONTROL_AUTH_VK_TOKEN_URL.
	ControlAuthVKTokenURL = env.GetEnvString(
		"CONTROL_AUTH_VK_TOKEN_URL",
		"",
	)

	// ControlAuthVKUserInfoURL contains the optional VK ID profile endpoint.
	// Env: CONTROL_AUTH_VK_USER_INFO_URL.
	ControlAuthVKUserInfoURL = env.GetEnvString(
		"CONTROL_AUTH_VK_USER_INFO_URL",
		"",
	)

	// ControlAuthGitHubClientID contains the GitHub OAuth client ID.
	// Env: CONTROL_AUTH_GITHUB_CLIENT_ID.
	ControlAuthGitHubClientID = env.GetEnvString(
		"CONTROL_AUTH_GITHUB_CLIENT_ID",
		"",
	)

	// ControlAuthGitHubClientSecret contains the GitHub OAuth client secret.
	// Env: CONTROL_AUTH_GITHUB_CLIENT_SECRET.
	ControlAuthGitHubClientSecret = env.GetEnvString(
		"CONTROL_AUTH_GITHUB_CLIENT_SECRET",
		"",
	)

	// ControlAuthGitHubTokenURL contains the optional GitHub OAuth token endpoint.
	// Env: CONTROL_AUTH_GITHUB_TOKEN_URL.
	ControlAuthGitHubTokenURL = env.GetEnvString(
		"CONTROL_AUTH_GITHUB_TOKEN_URL",
		"",
	)

	// ControlAuthGitHubUserInfoURL contains the optional GitHub profile endpoint.
	// Env: CONTROL_AUTH_GITHUB_USER_INFO_URL.
	ControlAuthGitHubUserInfoURL = env.GetEnvString(
		"CONTROL_AUTH_GITHUB_USER_INFO_URL",
		"",
	)

	// ControlAuthGitLabClientID contains the GitLab OAuth client ID.
	// Env: CONTROL_AUTH_GITLAB_CLIENT_ID.
	ControlAuthGitLabClientID = env.GetEnvString(
		"CONTROL_AUTH_GITLAB_CLIENT_ID",
		"",
	)

	// ControlAuthGitLabClientSecret contains the GitLab OAuth client secret.
	// Env: CONTROL_AUTH_GITLAB_CLIENT_SECRET.
	ControlAuthGitLabClientSecret = env.GetEnvString(
		"CONTROL_AUTH_GITLAB_CLIENT_SECRET",
		"",
	)

	// ControlAuthGitLabTokenURL contains the optional GitLab OAuth token endpoint.
	// Env: CONTROL_AUTH_GITLAB_TOKEN_URL.
	ControlAuthGitLabTokenURL = env.GetEnvString(
		"CONTROL_AUTH_GITLAB_TOKEN_URL",
		"",
	)

	// ControlAuthGitLabUserInfoURL contains the optional GitLab profile endpoint.
	// Env: CONTROL_AUTH_GITLAB_USER_INFO_URL.
	ControlAuthGitLabUserInfoURL = env.GetEnvString(
		"CONTROL_AUTH_GITLAB_USER_INFO_URL",
		"",
	)

	// ControlAuthGoogleClientID contains the Google OAuth client ID.
	// Env: CONTROL_AUTH_GOOGLE_CLIENT_ID.
	ControlAuthGoogleClientID = env.GetEnvString(
		"CONTROL_AUTH_GOOGLE_CLIENT_ID",
		"",
	)

	// ControlAuthGoogleClientSecret contains the Google OAuth client secret.
	// Env: CONTROL_AUTH_GOOGLE_CLIENT_SECRET.
	ControlAuthGoogleClientSecret = env.GetEnvString(
		"CONTROL_AUTH_GOOGLE_CLIENT_SECRET",
		"",
	)

	// ControlAuthGoogleTokenURL contains the optional Google OAuth token endpoint.
	// Env: CONTROL_AUTH_GOOGLE_TOKEN_URL.
	ControlAuthGoogleTokenURL = env.GetEnvString(
		"CONTROL_AUTH_GOOGLE_TOKEN_URL",
		"",
	)

	// ControlAuthGoogleUserInfoURL contains the optional Google profile endpoint.
	// Env: CONTROL_AUTH_GOOGLE_USER_INFO_URL.
	ControlAuthGoogleUserInfoURL = env.GetEnvString(
		"CONTROL_AUTH_GOOGLE_USER_INFO_URL",
		"",
	)

	// ControlAuthYandexClientID contains the Yandex OAuth client ID.
	// Env: CONTROL_AUTH_YANDEX_CLIENT_ID.
	ControlAuthYandexClientID = env.GetEnvString(
		"CONTROL_AUTH_YANDEX_CLIENT_ID",
		"",
	)

	// ControlAuthYandexClientSecret contains the Yandex OAuth client secret.
	// Env: CONTROL_AUTH_YANDEX_CLIENT_SECRET.
	ControlAuthYandexClientSecret = env.GetEnvString(
		"CONTROL_AUTH_YANDEX_CLIENT_SECRET",
		"",
	)

	// ControlAuthYandexTokenURL contains the optional Yandex OAuth token endpoint.
	// Env: CONTROL_AUTH_YANDEX_TOKEN_URL.
	ControlAuthYandexTokenURL = env.GetEnvString(
		"CONTROL_AUTH_YANDEX_TOKEN_URL",
		"",
	)

	// ControlAuthYandexUserInfoURL contains the optional Yandex profile endpoint.
	// Env: CONTROL_AUTH_YANDEX_USER_INFO_URL.
	ControlAuthYandexUserInfoURL = env.GetEnvString(
		"CONTROL_AUTH_YANDEX_USER_INFO_URL",
		"",
	)

	// ControlAuthTelegramBotToken contains the Telegram WebApp bot token.
	// Env: CONTROL_AUTH_TELEGRAM_BOT_TOKEN.
	ControlAuthTelegramBotToken = env.GetEnvString(
		"CONTROL_AUTH_TELEGRAM_BOT_TOKEN",
		"",
	)

	// ControlAuthTelegramMaxAge contains the maximum Telegram WebApp init data age.
	// Env: CONTROL_AUTH_TELEGRAM_MAX_AGE.
	ControlAuthTelegramMaxAge = env.GetEnvDuration(
		"CONTROL_AUTH_TELEGRAM_MAX_AGE",
		5*time.Minute,
	)

	// ControlAuthTONPayloadSecret contains the TON Connect payload signing secret.
	// Env: CONTROL_AUTH_TON_PAYLOAD_SECRET.
	ControlAuthTONPayloadSecret = env.GetEnvString(
		"CONTROL_AUTH_TON_PAYLOAD_SECRET",
		"",
	)

	// ControlAuthTONDomain contains the expected TON Connect proof domain.
	// Env: CONTROL_AUTH_TON_DOMAIN.
	ControlAuthTONDomain = env.GetEnvString(
		"CONTROL_AUTH_TON_DOMAIN",
		"",
	)

	// ControlAuthTONNetwork contains the optional expected TON network.
	// Env: CONTROL_AUTH_TON_NETWORK.
	ControlAuthTONNetwork = env.GetEnvString(
		"CONTROL_AUTH_TON_NETWORK",
		"",
	)

	// ControlAuthTONAllowNativeDomain contains the TON native-domain allowance.
	// Env: CONTROL_AUTH_TON_ALLOW_NATIVE_DOMAIN.
	ControlAuthTONAllowNativeDomain = env.GetEnvBool(
		"CONTROL_AUTH_TON_ALLOW_NATIVE_DOMAIN",
		false,
	)

	// ControlAuthTONMaxAge contains the maximum TON Connect proof age.
	// Env: CONTROL_AUTH_TON_MAX_AGE.
	ControlAuthTONMaxAge = env.GetEnvDuration(
		"CONTROL_AUTH_TON_MAX_AGE",
		5*time.Minute,
	)
)
