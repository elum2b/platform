package tasks

import (
	tcondition "github.com/elum2b/platform/internal/api/methods/tasks/admin/complex/condition"
	texport "github.com/elum2b/platform/internal/api/methods/tasks/admin/export"
	tgroup "github.com/elum2b/platform/internal/api/methods/tasks/admin/group"
	tgrouplocalization "github.com/elum2b/platform/internal/api/methods/tasks/admin/group/localization"
	timport "github.com/elum2b/platform/internal/api/methods/tasks/admin/import"
	tpartnerconfig "github.com/elum2b/platform/internal/api/methods/tasks/admin/partner/config"
	tpartnerdailystats "github.com/elum2b/platform/internal/api/methods/tasks/admin/partner/daily/stats"
	tpartnerrewardrule "github.com/elum2b/platform/internal/api/methods/tasks/admin/partner/reward/rule"
	treward "github.com/elum2b/platform/internal/api/methods/tasks/admin/reward"
	tsequence "github.com/elum2b/platform/internal/api/methods/tasks/admin/sequence"
	tstats "github.com/elum2b/platform/internal/api/methods/tasks/admin/stats"
	tstatsdaily "github.com/elum2b/platform/internal/api/methods/tasks/admin/stats/daily"
	tstatstask "github.com/elum2b/platform/internal/api/methods/tasks/admin/stats/task"
	ttask "github.com/elum2b/platform/internal/api/methods/tasks/admin/task"
	ttasklocalization "github.com/elum2b/platform/internal/api/methods/tasks/admin/task/localization"
	tuser "github.com/elum2b/platform/internal/api/methods/tasks/user"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all Tasks API methods.
func Register(registry adapter.Registry) {

	admin := registry.Group(adapter.AccountRequired)

	ttask.Save.Register(admin)
	ttask.Delete.Register(admin)
	ttask.Get.Register(admin)
	ttask.List.Register(admin)
	ttasklocalization.Upsert.Register(admin)
	tgroup.Upsert.Register(admin)
	tgrouplocalization.Upsert.Register(admin)
	tsequence.Upsert.Register(admin)
	treward.Upsert.Register(admin)
	treward.Delete.Register(admin)
	tcondition.Upsert.Register(admin)
	tcondition.Delete.Register(admin)
	tcondition.List.Register(admin)
	tpartnerconfig.Save.Register(admin)
	tpartnerconfig.Get.Register(admin)
	tpartnerconfig.List.Register(admin)
	tpartnerrewardrule.Save.Register(admin)
	tpartnerrewardrule.Delete.Register(admin)
	tpartnerdailystats.List.Register(admin)
	tstats.Get.Register(admin)
	tstatstask.TaskGet.Register(admin)
	tstatsdaily.List.Register(admin)
	tstatsdaily.Overview.Register(admin)
	tstatsdaily.Refresh.Register(admin)
	texport.Manifest.Register(admin)
	texport.Method.Register(admin)
	timport.Preview.Register(admin)
	timport.Method.Register(admin)

	user := registry.Group(adapter.ApplicationUser)

	tuser.ListActive.Register(user)
	tuser.Start.Register(user)
	tuser.Claim.Register(user)
	tuser.PartnerList.Register(user)
	tuser.PartnerCheck.Register(user)
	tuser.PartnerStart.Register(user)

}
