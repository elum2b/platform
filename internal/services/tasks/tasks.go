package tasks

import (
	"context"
	"path/filepath"

	"github.com/elum2b/services/tasks"
	taskruntime "github.com/elum2b/services/tasks/runtime"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"
)

func Service() func(context.Context) error {
	return func(ctx context.Context) error {
		if err := services.Tasks.OnCallback(ctx, handler); err != nil {
			return err
		}

		return services.Tasks.Run(ctx, tasks.DatabaseParams{
			Host:     config.TasksPostgresHost,
			Port:     config.TasksPostgresPort,
			User:     config.TasksPostgresUser,
			Password: config.TasksPostgresPassword,
			Database: config.TasksPostgresDatabase,
			SSLMode:  "disable",
			Options: tasks.Options{
				MaxConnections: config.TasksMaxConnections,
				QueryTimeout:   config.TasksQueryTimeout,
				CacheL1Delay:   config.TasksCacheL1Delay,
				CacheL2Delay:   config.TasksCacheL2Delay,
				CacheEnabled:   config.TasksCacheEnabled,
				CacheSize:      config.TasksCacheSize,
				CacheTTLCheck:  config.TasksCacheTTLCheck,
				ArchiveDirectory: filepath.Join(
					config.ServicesDataDirectory,
					"tasks",
				),

				PartnerStartLeaseDuration: config.TasksPartnerStartLeaseDuration,
				Runtime: taskruntime.Options{
					Enabled:          config.TasksRuntimeEnabled,
					ScriptCacheTTL:   config.TasksRuntimeScriptCacheTTL,
					Timeout:          config.TasksRuntimeTimeout,
					MaxMemory:        config.TasksRuntimeMaxMemory,
					MaxHTTPRequests:  config.TasksRuntimeMaxHTTPRequests,
					MaxResponseBytes: config.TasksRuntimeMaxResponseBytes,
					JSONBoundary:     config.TasksRuntimeJSONBoundary,
					StatePoolSize:    config.TasksRuntimeStatePoolSize,
				},
			},
		})
	}
}
