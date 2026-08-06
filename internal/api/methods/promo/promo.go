package promo

import (
	promocallback "github.com/elum2b/platform/internal/api/methods/promo/admin/callback"
	promoexport "github.com/elum2b/platform/internal/api/methods/promo/admin/export"
	promoimport "github.com/elum2b/platform/internal/api/methods/promo/admin/import"
	promolocalization "github.com/elum2b/platform/internal/api/methods/promo/admin/localization"
	promopromo "github.com/elum2b/platform/internal/api/methods/promo/admin/promo"
	promoredemption "github.com/elum2b/platform/internal/api/methods/promo/admin/redemption"
	promoreward "github.com/elum2b/platform/internal/api/methods/promo/admin/reward"
	promostats "github.com/elum2b/platform/internal/api/methods/promo/admin/stats"
	promoapply "github.com/elum2b/platform/internal/api/methods/promo/user"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all Promo API methods.
func Register(registry adapter.Registry) {
	admin := registry.Group(adapter.AccountRequired)

	promopromo.Upsert.Register(admin)
	promopromo.Get.Register(admin)
	promopromo.List.Register(admin)
	promopromo.Delete.Register(admin)
	promolocalization.Upsert.Register(admin)
	promolocalization.Get.Register(admin)
	promolocalization.List.Register(admin)
	promolocalization.Delete.Register(admin)
	promoreward.Upsert.Register(admin)
	promoreward.Get.Register(admin)
	promoreward.List.Register(admin)
	promoreward.Delete.Register(admin)
	promoexport.Method.Register(admin)
	promoimport.Preview.Register(admin)
	promoimport.Method.Register(admin)
	promostats.Get.Register(admin)
	promostats.DailyList.Register(admin)
	promostats.Refresh.Register(admin)
	promoredemption.Get.Register(admin)
	promoredemption.List.Register(admin)
	promocallback.List.Register(admin)
	promocallback.Get.Register(admin)
	promocallback.Retry.Register(admin)
	promocallback.MarkOK.Register(admin)
	promocallback.MarkReject.Register(admin)
	promocallback.ResetExpired.Register(admin)

	user := registry.Group(adapter.ApplicationUser)

	promoapply.Method.Register(user)
}
