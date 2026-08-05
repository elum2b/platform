package calendar

import (
	calcalendar "github.com/elum2b/platform/internal/api/methods/calendar/admin/calendar"
	calcallback "github.com/elum2b/platform/internal/api/methods/calendar/admin/callback"
	calexport "github.com/elum2b/platform/internal/api/methods/calendar/admin/export"
	calimport "github.com/elum2b/platform/internal/api/methods/calendar/admin/import"
	callocalization "github.com/elum2b/platform/internal/api/methods/calendar/admin/localization"
	caloperation "github.com/elum2b/platform/internal/api/methods/calendar/admin/operation"
	calreward "github.com/elum2b/platform/internal/api/methods/calendar/admin/reward"
	calstats "github.com/elum2b/platform/internal/api/methods/calendar/admin/stats"
	calstep "github.com/elum2b/platform/internal/api/methods/calendar/admin/step"
	caluser "github.com/elum2b/platform/internal/api/methods/calendar/user"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all Calendar API methods.
func Register(registry adapter.Registry) {

	admin := registry.Group(adapter.AccountRequired)

	calcalendar.Upsert.Register(admin)
	calcalendar.Get.Register(admin)
	calcalendar.List.Register(admin)
	calcalendar.SetActive.Register(admin)
	calcalendar.Delete.Register(admin)
	callocalization.Upsert.Register(admin)
	callocalization.Get.Register(admin)
	callocalization.List.Register(admin)
	callocalization.Delete.Register(admin)
	calstep.Create.Register(admin)
	calstep.Update.Register(admin)
	calstep.Delete.Register(admin)
	calreward.Create.Register(admin)
	calreward.Update.Register(admin)
	calreward.Get.Register(admin)
	calreward.Delete.Register(admin)
	caloperation.List.Register(admin)
	calstats.Get.Register(admin)
	calstats.DailyList.Register(admin)
	calstats.Refresh.Register(admin)
	calexport.Method.Register(admin)
	calimport.Preview.Register(admin)
	calimport.Method.Register(admin)
	calcallback.List.Register(admin)
	calcallback.Get.Register(admin)
	calcallback.Retry.Register(admin)
	calcallback.MarkOK.Register(admin)
	calcallback.MarkReject.Register(admin)
	calcallback.ResetExpired.Register(admin)

	user := registry.Group(adapter.ApplicationUser)

	caluser.ListActive.Register(user)
	caluser.Get.Register(user)
	caluser.GetProgress.Register(user)
	caluser.Next.Register(user)
	caluser.Record.Register(user)

}
