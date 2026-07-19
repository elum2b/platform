package tasks

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"

	"github.com/elum2b/services/tasks"
	taskruntime "github.com/elum2b/services/tasks/runtime"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		services.Tasks = tasks.New(tasks.DatabaseParams{
			Host:     config.TasksPostgresHost,
			Port:     config.TasksPostgresPort,
			User:     config.TasksPostgresUser,
			Password: config.TasksPostgresPassword,
			Database: config.TasksPostgresDatabase,
			Options: tasks.Options{
				MaxConnections: config.TasksMaxConnections,
				QueryTimeout:   config.TasksQueryTimeout,
				CacheL1Delay:   config.TasksCacheL1Delay,
				CacheL2Delay:   config.TasksCacheL2Delay,
				CacheEnabled:   config.TasksCacheEnabled,
				CacheSize:      config.TasksCacheSize,
				CacheTTLCheck:  config.TasksCacheTTLCheck,

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

		if err := services.Tasks.OnCallback(ctx, handler); err != nil {
			return err
		}

		return services.Tasks.Run(ctx)

	}

}
