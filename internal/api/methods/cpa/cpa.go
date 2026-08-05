package cpa

import (
	cpaassignment "github.com/elum2b/platform/internal/api/methods/cpa/admin/assignment"
	cpaevent "github.com/elum2b/platform/internal/api/methods/cpa/admin/assignment/event"
	cpacallback "github.com/elum2b/platform/internal/api/methods/cpa/admin/callback"
	cpacode "github.com/elum2b/platform/internal/api/methods/cpa/admin/code"
	cpaexport "github.com/elum2b/platform/internal/api/methods/cpa/admin/export"
	cpaimport "github.com/elum2b/platform/internal/api/methods/cpa/admin/import"
	cpalocalization "github.com/elum2b/platform/internal/api/methods/cpa/admin/localization"
	cpaoffer "github.com/elum2b/platform/internal/api/methods/cpa/admin/offer"
	cpareward "github.com/elum2b/platform/internal/api/methods/cpa/admin/reward"
	cpastats "github.com/elum2b/platform/internal/api/methods/cpa/admin/stats"
	usercode "github.com/elum2b/platform/internal/api/methods/cpa/user/code"
	useroffer "github.com/elum2b/platform/internal/api/methods/cpa/user/offer"
	userstatus "github.com/elum2b/platform/internal/api/methods/cpa/user/status"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all CPA API methods.
func Register(registry adapter.Registry) {

	admin := registry.Group(adapter.AccountRequired)

	cpaoffer.Upsert.Register(admin)
	cpaoffer.Get.Register(admin)
	cpaoffer.List.Register(admin)
	cpaoffer.Delete.Register(admin)
	cpalocalization.Upsert.Register(admin)
	cpalocalization.List.Register(admin)
	cpalocalization.Delete.Register(admin)
	cpareward.Upsert.Register(admin)
	cpareward.List.Register(admin)
	cpareward.Delete.Register(admin)
	cpaexport.Method.Register(admin)
	cpaimport.Preview.Register(admin)
	cpaimport.Method.Register(admin)
	cpacode.Add.Register(admin)
	cpacode.DeleteAvailable.Register(admin)
	cpacode.DeleteIssued.Register(admin)
	cpacode.DeleteCompleted.Register(admin)
	cpaassignment.Get.Register(admin)
	cpaassignment.List.Register(admin)
	cpacode.List.Register(admin)
	cpaevent.List.Register(admin)
	cpaassignment.Complete.Register(admin)
	cpastats.Get.Register(admin)
	cpastats.DailyList.Register(admin)
	cpastats.Refresh.Register(admin)
	cpacallback.List.Register(admin)
	cpacallback.Get.Register(admin)
	cpacallback.Retry.Register(admin)
	cpacallback.MarkOK.Register(admin)
	cpacallback.MarkReject.Register(admin)
	cpacallback.ResetExpired.Register(admin)

	user := registry.Group(adapter.ApplicationUser)

	useroffer.List.Register(user)
	usercode.Get.Register(user)
	userstatus.Get.Register(user)

}
