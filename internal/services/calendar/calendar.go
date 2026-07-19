package calendar

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/services"

	"github.com/elum2b/services/calendar"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		services.Calendar = calendar.New(calendar.DatabaseParams{
			Host:     config.CalendarPostgresHost,
			Port:     config.CalendarPostgresPort,
			User:     config.CalendarPostgresUser,
			Password: config.CalendarPostgresPassword,
			Database: config.CalendarPostgresDatabase,
			Options: calendar.Options{
				MaxConnections: config.CalendarMaxConnections,
				QueryTimeout:   config.CalendarQueryTimeout,
				CacheL1Delay:   config.CalendarCacheL1Delay,
				CacheL2Delay:   config.CalendarCacheL2Delay,
				CacheEnabled:   config.CalendarCacheEnabled,
				CacheSize:      config.CalendarCacheSize,
				CacheTTLCheck:  config.CalendarCacheTTLCheck,
			},
		})

		if err := services.Calendar.OnCallback(ctx, handler); err != nil {
			return err
		}

		return services.Calendar.Run(ctx)

	}

}
