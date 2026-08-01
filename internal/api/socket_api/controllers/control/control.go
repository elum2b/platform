package control

import (
	etp "github.com/elum-utils/go-etp"

	"github.com/elum2b/platform/internal/api/socket_api/controllers/control/account"
	accountaccess "github.com/elum2b/platform/internal/api/socket_api/controllers/control/account/access"
	accountidentity "github.com/elum2b/platform/internal/api/socket_api/controllers/control/account/identity"
	accountlimitrequest "github.com/elum2b/platform/internal/api/socket_api/controllers/control/account/limit_request"
	accountsession "github.com/elum2b/platform/internal/api/socket_api/controllers/control/account/session"
	accounttwofactor "github.com/elum2b/platform/internal/api/socket_api/controllers/control/account/two_factor"
	"github.com/elum2b/platform/internal/api/socket_api/controllers/control/global"
	globalaccess "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/access"
	globalaudit "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/audit"
	globalinvite "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/invite"
	globallimit "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/limit"
	globalmember "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/member"
	globalmethod "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/method"
	globalrole "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/role"
	globalrolemember "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/role/member"
	globalperm "github.com/elum2b/platform/internal/api/socket_api/controllers/control/global/role/permission"
	"github.com/elum2b/platform/internal/api/socket_api/controllers/control/workspace"
	workspaceaudit "github.com/elum2b/platform/internal/api/socket_api/controllers/control/workspace/audit"
	workspaceinvite "github.com/elum2b/platform/internal/api/socket_api/controllers/control/workspace/invite"
	workspacerole "github.com/elum2b/platform/internal/api/socket_api/controllers/control/workspace/role"
	workspacerolemember "github.com/elum2b/platform/internal/api/socket_api/controllers/control/workspace/role/member"
	workspaceperm "github.com/elum2b/platform/internal/api/socket_api/controllers/control/workspace/role/permission"
	"github.com/elum2b/platform/internal/api/socket_api/middleware"
)

// Register registers all control socket handlers.
func Register(socket etp.Router) {
	group := socket.Group()

	group.Use("*", middleware.Authenticated)
	group.Use("*", middleware.ControlReady)

	account.Get("control.account.get", group)
	accountidentity.List("control.account.identity.list", group)
	accountidentity.Bind("control.account.identity.bind", group)
	accountidentity.Unbind("control.account.identity.unbind", group)
	accountsession.List("control.account.session.list", group)
	accountsession.Revoke("control.account.session.revoke", group)
	accountsession.RevokeAll("control.account.session.revoke_all", group)
	accounttwofactor.Begin("control.account.two_factor.begin", group)
	accounttwofactor.Confirm("control.account.two_factor.confirm", group)
	accounttwofactor.Disable("control.account.two_factor.disable", group)
	accountlimitrequest.WorkspaceCreate("control.account.workspace_limit.request", group)
	accountlimitrequest.Cancel("control.account.limit_request.cancel", group)
	accountaccess.GlobalList("control.account.access.global.list", group)
	accountaccess.WorkspaceList("control.account.access.workspace.list", group)

	workspace.Create("control.global.workspace.create", group)
	workspace.Get("control.workspace.get", group)
	workspace.List("control.workspace.list", group)
	workspace.Update("control.workspace.update", group)
	workspace.Archive("control.workspace.archive", group)
	workspace.OwnerTransfer("control.workspace.owner.transfer", group)
	workspace.MemberList("control.workspace.member.list", group)
	workspace.MemberRemove("control.workspace.member.remove", group)
	workspace.EmployeeLimitRequest("control.workspace.employee_limit.request", group)
	workspacerole.Create("control.workspace.role.create", group)
	workspacerole.List("control.workspace.role.list", group)
	workspacerole.Get("control.workspace.role.get", group)
	workspacerole.Update("control.workspace.role.update", group)
	workspacerole.Delete("control.workspace.role.delete", group)
	workspacerolemember.Assign("control.workspace.role.member.assign", group)
	workspacerolemember.Remove("control.workspace.role.member.remove", group)
	workspaceperm.Replace("control.workspace.role.permission.replace", group)
	workspaceperm.List("control.workspace.role.permission.list", group)
	workspaceinvite.Create("control.workspace.invite.create", group)
	workspaceinvite.List("control.workspace.invite.list", group)
	workspaceinvite.Revoke("control.workspace.invite.revoke", group)
	workspaceaudit.List("control.workspace.audit.list", group)

	global.OwnerTransfer("control.global.owner.transfer", group)
	globalmember.List("control.global.member.list", group)
	globalmember.Remove("control.global.member.remove", group)
	globalrole.Create("control.global.role.create", group)
	globalrole.List("control.global.role.list", group)
	globalrole.Get("control.global.role.get", group)
	globalrole.Update("control.global.role.update", group)
	globalrole.Delete("control.global.role.delete", group)
	globalrolemember.Assign("control.global.role.member.assign", group)
	globalrolemember.Remove("control.global.role.member.remove", group)
	globalperm.Replace("control.global.role.permission.replace", group)
	globalperm.List("control.global.role.permission.list", group)
	globalinvite.Create("control.global.invite.create", group)
	globalinvite.List("control.global.invite.list", group)
	globalinvite.Revoke("control.global.invite.revoke", group)
	globalmethod.List("control.global.method.list", group)
	globalmethod.Get("control.global.method.get", group)
	globalaccess.List("control.global.access.list", group)
	globallimit.List("control.global.limit.list", group)
	globallimit.Resolve("control.global.limit.resolve", group)
	globalaudit.List("control.global.audit.list", group)
}
