package methods

import (
	controlaccount "github.com/elum2b/platform/internal/api/methods/control/account"
	accountaccess "github.com/elum2b/platform/internal/api/methods/control/account/access"
	accountidentity "github.com/elum2b/platform/internal/api/methods/control/account/identity"
	accountlimit "github.com/elum2b/platform/internal/api/methods/control/account/limit_request"
	accountmcp "github.com/elum2b/platform/internal/api/methods/control/account/mcp"
	accountsession "github.com/elum2b/platform/internal/api/methods/control/account/session"
	accounttwofactor "github.com/elum2b/platform/internal/api/methods/control/account/two_factor"
	controlauth "github.com/elum2b/platform/internal/api/methods/control/auth"
	controlglobal "github.com/elum2b/platform/internal/api/methods/control/global"
	globalaccess "github.com/elum2b/platform/internal/api/methods/control/global/access"
	globalaudit "github.com/elum2b/platform/internal/api/methods/control/global/audit"
	globalinvite "github.com/elum2b/platform/internal/api/methods/control/global/invite"
	globallimit "github.com/elum2b/platform/internal/api/methods/control/global/limit"
	globalmember "github.com/elum2b/platform/internal/api/methods/control/global/member"
	globalmethod "github.com/elum2b/platform/internal/api/methods/control/global/method"
	globalrole "github.com/elum2b/platform/internal/api/methods/control/global/role"
	globalrolemember "github.com/elum2b/platform/internal/api/methods/control/global/role/member"
	globalrolepermission "github.com/elum2b/platform/internal/api/methods/control/global/role/permission"
	controlworkspace "github.com/elum2b/platform/internal/api/methods/control/workspace"
	workspaceaudit "github.com/elum2b/platform/internal/api/methods/control/workspace/audit"
	workspaceinvite "github.com/elum2b/platform/internal/api/methods/control/workspace/invite"
	workspacerole "github.com/elum2b/platform/internal/api/methods/control/workspace/role"
	workspacerolemember "github.com/elum2b/platform/internal/api/methods/control/workspace/role/member"
	workspacerolepermission "github.com/elum2b/platform/internal/api/methods/control/workspace/role/permission"
	"github.com/elum2b/platform/internal/api/methods/cpa/offer"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all platform API methods.
func Register(registry adapter.Registry) {
	controlauth.Check.Register(registry)
	controlauth.VKID.Register(registry)
	controlauth.Telegram.Register(registry)
	controlauth.GitHub.Register(registry)
	controlauth.GitLab.Register(registry)
	controlauth.Google.Register(registry)
	controlauth.Yandex.Register(registry)
	controlauth.TON.Register(registry)
	controlauth.TONChallenge.Register(registry)
	controlauth.TwoFactor.Register(registry)
	controlaccount.Get.Register(registry)
	accountidentity.List.Register(registry)
	accountidentity.Bind.Register(registry)
	accountidentity.Unbind.Register(registry)
	accountsession.List.Register(registry)
	accountsession.Revoke.Register(registry)
	accountsession.RevokeAll.Register(registry)
	accountmcp.Create.Register(registry)
	accountmcp.List.Register(registry)
	accountmcp.Revoke.Register(registry)
	accounttwofactor.Begin.Register(registry)
	accounttwofactor.Confirm.Register(registry)
	accounttwofactor.Disable.Register(registry)
	accountlimit.WorkspaceCreate.Register(registry)
	accountlimit.Cancel.Register(registry)
	accountaccess.GlobalList.Register(registry)
	accountaccess.WorkspaceList.Register(registry)
	offer.List.Register(registry)
	controlworkspace.Get.Register(registry)
	controlworkspace.Create.Register(registry)
	controlworkspace.List.Register(registry)
	controlworkspace.Update.Register(registry)
	controlworkspace.Archive.Register(registry)
	controlworkspace.OwnerTransfer.Register(registry)
	controlworkspace.MemberRemove.Register(registry)
	controlworkspace.MemberList.Register(registry)
	controlworkspace.EmployeeLimit.Register(registry)
	workspaceaudit.List.Register(registry)
	workspaceinvite.Create.Register(registry)
	workspaceinvite.List.Register(registry)
	workspaceinvite.Revoke.Register(registry)
	workspacerole.Create.Register(registry)
	workspacerole.List.Register(registry)
	workspacerole.Get.Register(registry)
	workspacerole.Update.Register(registry)
	workspacerole.Delete.Register(registry)
	workspacerolemember.Assign.Register(registry)
	workspacerolemember.Remove.Register(registry)
	workspacerolepermission.Replace.Register(registry)
	workspacerolepermission.List.Register(registry)
	controlglobal.OwnerTransfer.Register(registry)
	globalmember.List.Register(registry)
	globalmember.Remove.Register(registry)
	globalrole.Create.Register(registry)
	globalrole.List.Register(registry)
	globalrole.Get.Register(registry)
	globalrole.Update.Register(registry)
	globalrole.Delete.Register(registry)
	globalrolemember.Assign.Register(registry)
	globalrolemember.Remove.Register(registry)
	globalrolepermission.Replace.Register(registry)
	globalrolepermission.List.Register(registry)
	globalinvite.Create.Register(registry)
	globalinvite.List.Register(registry)
	globalinvite.Revoke.Register(registry)
	globalmethod.List.Register(registry)
	globalmethod.Get.Register(registry)
	globalaccess.List.Register(registry)
	globallimit.List.Register(registry)
	globallimit.Resolve.Register(registry)
	globalaudit.List.Register(registry)
}
