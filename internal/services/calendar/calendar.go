package calendar

import (
	"context"

	"github.com/elum2b/platform/internal/config"
	"github.com/elum2b/platform/internal/global"
	service "github.com/elum2b/services/calendar"
)

func Service() func(context.Context) error {

	return func(ctx context.Context) error {

		global.Calendar = service.New(service.DatabaseParams{
			Host:     config.CalendarPostgresHost,
			Port:     config.CalendarPostgresPort,
			User:     config.CalendarPostgresUser,
			Password: config.CalendarPostgresPassword,
			Database: config.CalendarPostgresDatabase,
			Options: service.Options{
				MaxConnections: config.CalendarMaxConnections,
				QueryTimeout:   config.CalendarQueryTimeout,
				CacheL1Delay:   config.CalendarCacheL1Delay,
				CacheL2Delay:   config.CalendarCacheL2Delay,
				CacheEnabled:   config.CalendarCacheEnabled,
				CacheSize:      config.CalendarCacheSize,
				CacheTTLCheck:  config.CalendarCacheTTLCheck,
			},
		})

		if err := global.Calendar.OnCallback(ctx, handler); err != nil {
			return err
		}

		return global.Calendar.Run(ctx)

	}

}
