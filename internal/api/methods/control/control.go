package control

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
	adapter "github.com/elum2b/platform/internal/utils/adapter"
)

// Register registers all Control API methods.
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

	account := registry.Group(adapter.AccountRequired)

	controlaccount.Get.Register(account)
	accountidentity.List.Register(account)
	accountidentity.Bind.Register(account)
	accountidentity.Unbind.Register(account)
	accountsession.List.Register(account)
	accountsession.Revoke.Register(account)
	accountsession.RevokeAll.Register(account)
	accountmcp.Create.Register(account)
	accountmcp.List.Register(account)
	accountmcp.Revoke.Register(account)
	accounttwofactor.Begin.Register(account)
	accounttwofactor.Confirm.Register(account)
	accounttwofactor.Disable.Register(account)
	accountlimit.WorkspaceCreate.Register(account)
	accountlimit.Cancel.Register(account)
	accountaccess.GlobalList.Register(account)
	accountaccess.WorkspaceList.Register(account)
	controlworkspace.Get.Register(account)
	controlworkspace.Create.Register(account)
	controlworkspace.List.Register(account)
	controlworkspace.Update.Register(account)
	controlworkspace.Archive.Register(account)
	controlworkspace.OwnerTransfer.Register(account)
	controlworkspace.MemberRemove.Register(account)
	controlworkspace.MemberList.Register(account)
	controlworkspace.EmployeeLimit.Register(account)
	workspaceaudit.List.Register(account)
	workspaceinvite.Create.Register(account)
	workspaceinvite.List.Register(account)
	workspaceinvite.Revoke.Register(account)
	workspacerole.Create.Register(account)
	workspacerole.List.Register(account)
	workspacerole.Get.Register(account)
	workspacerole.Update.Register(account)
	workspacerole.Delete.Register(account)
	workspacerolemember.Assign.Register(account)
	workspacerolemember.Remove.Register(account)
	workspacerolepermission.Replace.Register(account)
	workspacerolepermission.List.Register(account)
	controlglobal.OwnerTransfer.Register(account)
	globalmember.List.Register(account)
	globalmember.Remove.Register(account)
	globalrole.Create.Register(account)
	globalrole.List.Register(account)
	globalrole.Get.Register(account)
	globalrole.Update.Register(account)
	globalrole.Delete.Register(account)
	globalrolemember.Assign.Register(account)
	globalrolemember.Remove.Register(account)
	globalrolepermission.Replace.Register(account)
	globalrolepermission.List.Register(account)
	globalinvite.Create.Register(account)
	globalinvite.List.Register(account)
	globalinvite.Revoke.Register(account)
	globalmethod.List.Register(account)
	globalmethod.Get.Register(account)
	globalaccess.List.Register(account)
	globallimit.List.Register(account)
	globallimit.Resolve.Register(account)
	globalaudit.List.Register(account)
}
