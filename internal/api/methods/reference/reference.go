package reference

import (
	refexport "github.com/elum2b/platform/internal/api/methods/reference/admin/export"
	refimport "github.com/elum2b/platform/internal/api/methods/reference/admin/import"
	refitem "github.com/elum2b/platform/internal/api/methods/reference/admin/item"
	reflocalization "github.com/elum2b/platform/internal/api/methods/reference/admin/localization"
	refstats "github.com/elum2b/platform/internal/api/methods/reference/admin/stats"
	refuser "github.com/elum2b/platform/internal/api/methods/reference/user"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all Reference API methods.
func Register(registry adapter.Registry) {

	admin := registry.Group(adapter.AccountRequired)

	refitem.Create.Register(admin)
	refitem.Update.Register(admin)
	refitem.ChangeType.Register(admin)
	refitem.Get.Register(admin)
	refitem.List.Register(admin)
	refitem.Delete.Register(admin)
	refitem.Restore.Register(admin)
	reflocalization.Upsert.Register(admin)
	reflocalization.Get.Register(admin)
	reflocalization.List.Register(admin)
	reflocalization.Delete.Register(admin)
	refstats.Get.Register(admin)
	refexport.Method.Register(admin)
	refimport.Preview.Register(admin)
	refimport.Method.Register(admin)

	user := registry.Group(adapter.ApplicationUser)

	refuser.Get.Register(user)
	refuser.Resolve.Register(user)
	refuser.List.Register(user)

}
