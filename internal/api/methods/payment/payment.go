package payment

import (
	payasset "github.com/elum2b/platform/internal/api/methods/payment/admin/asset"
	payassetrate "github.com/elum2b/platform/internal/api/methods/payment/admin/asset/rate"
	paycallback "github.com/elum2b/platform/internal/api/methods/payment/admin/callback"
	paycatalog "github.com/elum2b/platform/internal/api/methods/payment/admin/catalog"
	paydailystats "github.com/elum2b/platform/internal/api/methods/payment/admin/daily/stats"
	payexport "github.com/elum2b/platform/internal/api/methods/payment/admin/export"
	payfulfillment "github.com/elum2b/platform/internal/api/methods/payment/admin/fulfillment"
	payfulfillmentitem "github.com/elum2b/platform/internal/api/methods/payment/admin/fulfillment/item"
	payimport "github.com/elum2b/platform/internal/api/methods/payment/admin/import"
	paylocalization "github.com/elum2b/platform/internal/api/methods/payment/admin/localization"
	payoperation "github.com/elum2b/platform/internal/api/methods/payment/admin/operation"
	payorder "github.com/elum2b/platform/internal/api/methods/payment/admin/order"
	payattempt "github.com/elum2b/platform/internal/api/methods/payment/admin/payment/attempt"
	payevent "github.com/elum2b/platform/internal/api/methods/payment/admin/payment/event"
	payprice "github.com/elum2b/platform/internal/api/methods/payment/admin/price"
	payproduct "github.com/elum2b/platform/internal/api/methods/payment/admin/product"
	payproductgroup "github.com/elum2b/platform/internal/api/methods/payment/admin/product/group"
	payproductitem "github.com/elum2b/platform/internal/api/methods/payment/admin/product/item"
	paylimitcounter "github.com/elum2b/platform/internal/api/methods/payment/admin/product/limit/counter"
	payprovider "github.com/elum2b/platform/internal/api/methods/payment/admin/provider"
	payproviderasset "github.com/elum2b/platform/internal/api/methods/payment/admin/provider/asset"
	payprovidercursor "github.com/elum2b/platform/internal/api/methods/payment/admin/provider/cursor"
	payprovidertx "github.com/elum2b/platform/internal/api/methods/payment/admin/provider/transaction"
	paypurchasekey "github.com/elum2b/platform/internal/api/methods/payment/admin/purchase/key"
	payrefund "github.com/elum2b/platform/internal/api/methods/payment/admin/refund"
	payreport "github.com/elum2b/platform/internal/api/methods/payment/admin/report"
	paysubscription "github.com/elum2b/platform/internal/api/methods/payment/admin/subscription"
	paytonwallet "github.com/elum2b/platform/internal/api/methods/payment/admin/ton/wallet"
	payuser "github.com/elum2b/platform/internal/api/methods/payment/user"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all Payment API methods.
func Register(registry adapter.Registry) {
	admin := registry.Group(adapter.AccountRequired)

	payproductgroup.Upsert.Register(admin)
	payproductgroup.Get.Register(admin)
	payproductgroup.List.Register(admin)
	payproductgroup.Delete.Register(admin)
	payproduct.Upsert.Register(admin)
	payproduct.Get.Register(admin)
	payproduct.List.Register(admin)
	payproduct.Delete.Register(admin)
	payproductitem.List.Register(admin)
	payproductitem.Upsert.Register(admin)
	payproductitem.Delete.Register(admin)
	payprice.List.Register(admin)
	payprice.Get.Register(admin)
	payprice.Create.Register(admin)
	payprice.Update.Register(admin)
	payprice.Delete.Register(admin)
	paylocalization.List.Register(admin)
	paylocalization.Get.Register(admin)
	paylocalization.Upsert.Register(admin)
	paylocalization.Delete.Register(admin)
	paylimitcounter.List.Register(admin)
	paylimitcounter.Delete.Register(admin)
	paycatalog.SaveProduct.Register(admin)
	paycatalog.SaveProductGroup.Register(admin)
	paycatalog.SaveLocalization.Register(admin)
	paycatalog.AttachProductItem.Register(admin)
	paycatalog.CreatePrice.Register(admin)
	paycatalog.UpdatePrice.Register(admin)
	payprovider.List.Register(admin)
	payprovider.Get.Register(admin)
	payasset.List.Register(admin)
	payasset.Get.Register(admin)
	payproviderasset.List.Register(admin)
	payproviderasset.Get.Register(admin)
	payassetrate.List.Register(admin)
	payassetrate.Get.Register(admin)
	paytonwallet.Save.Register(admin)
	paytonwallet.Get.Register(admin)
	paytonwallet.Delete.Register(admin)
	payoperation.CreateProductKey.Register(admin)
	payoperation.RebuildProductCache.Register(admin)
	payoperation.ExecuteRefund.Register(admin)
	payexport.Method.Register(admin)
	payimport.Preview.Register(admin)
	payimport.Method.Register(admin)
	paypurchasekey.List.Register(admin)
	paypurchasekey.Get.Register(admin)
	paypurchasekey.UpdateStatus.Register(admin)
	payorder.List.Register(admin)
	payorder.ListUser.Register(admin)
	payorder.Get.Register(admin)
	payorder.GetByPublicID.Register(admin)
	payorder.UpdateStatus.Register(admin)
	payattempt.List.Register(admin)
	payattempt.Get.Register(admin)
	payattempt.UpdateStatus.Register(admin)
	payevent.List.Register(admin)
	payevent.Get.Register(admin)
	payevent.UpdateProcessingStatus.Register(admin)
	paysubscription.List.Register(admin)
	paysubscription.Get.Register(admin)
	paysubscription.GetByProviderID.Register(admin)
	paysubscription.Upsert.Register(admin)
	paysubscription.UpdateStatus.Register(admin)
	payfulfillment.List.Register(admin)
	payfulfillment.Get.Register(admin)
	payfulfillment.UpdateStatus.Register(admin)
	payfulfillmentitem.List.Register(admin)
	payrefund.Create.Register(admin)
	payrefund.List.Register(admin)
	payrefund.Get.Register(admin)
	payrefund.UpdateStatus.Register(admin)
	payreport.Get.Register(admin)
	paydailystats.List.Register(admin)
	paydailystats.Overview.Register(admin)
	paydailystats.Refresh.Register(admin)
	payprovidercursor.List.Register(admin)
	payprovidercursor.Get.Register(admin)
	payprovidercursor.Upsert.Register(admin)
	payprovidertx.List.Register(admin)
	payprovidertx.Get.Register(admin)
	payprovidertx.GetByExternalID.Register(admin)
	payprovidertx.UpdateStatus.Register(admin)
	paycallback.List.Register(admin)
	paycallback.Get.Register(admin)
	paycallback.Retry.Register(admin)
	paycallback.MarkOK.Register(admin)
	paycallback.MarkReject.Register(admin)
	paycallback.ResetExpired.Register(admin)

	user := registry.Group(adapter.ApplicationUser)

	payuser.ListProducts.Register(user)
	payuser.GetProduct.Register(user)
	payuser.GetProductByKey.Register(user)
	payuser.ListAssets.Register(user)
	payuser.GetUSDTPrice.Register(user)
	payuser.ListUSDTPrices.Register(user)
	payuser.CreateOrder.Register(user)
	payuser.CreateOrderByKey.Register(user)
	payuser.CreateAttempt.Register(user)
	payuser.IsSubscriptionActive.Register(user)
}
