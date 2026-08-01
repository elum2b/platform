# Platform backend

## Purpose and boundaries

- This repository is an external API wrapper around
  `github.com/elum2b/services`. Business rules, persistence and migrations
  belong to that dependency; do not access PostgreSQL, Redis or another store
  directly from this repository.
- Service lifecycle belongs to `internal/services/<service>`. Register service
  runners in `main.go` through `github.com/elum-utils/supervisor` and preserve
  the existing explanatory comment style there.
- All environment variables live in `internal/config`. Use only
  `github.com/elum-utils/env`; do not add configuration helper functions.
  Comments are two lines: a concise description and `Env:` or `Fallback:`.

## HTTP authentication

- Fiber controllers are under `internal/api/internal_api`.
- HTTP is reserved for authentication: provider login, two-factor completion,
  session check and TON challenge. Account logout and other authenticated work
  are socket methods.
- Keep session tokens only in cookies through `internal/utils/cookies`.
  Cookies must remain `Secure`, `HTTPOnly` and configured with `SameSite`.
- Provider-specific identity resolution belongs in `internal/utils/auth`, not
  in a controller.

## Socket API

- All non-authentication API methods use ETP WebSocket controllers under
  `internal/api/socket_api/controllers`.
- Register all control events only in
  `internal/api/socket_api/controllers/control/control.go`. `socket.go` calls
  only `control.Register(ws)`.
- Register a controller with this shape:

  ```go
  func Method(event string, socket etp.Router) {
      socket.On(event, func(ctx *etp.Context) error {
          // decode, call services, respond
      })
  }
  ```

- Controllers are thin: decode via `internal/utils/socket`, call
  `internal/services`, map a response, then respond via the same utility.
  Return service errors unchanged.
- One file contains one handler. Keep request and response types in that same
  file; never create a shared `models.go` for a controller section.
- Use nested packages for controller domains, for example
  `account/identity/bind.go` or `workspace/role/permission/replace.go`.
- Add `middleware.GlobalAccess` or `middleware.WorkspaceAccess` using the exact
  service method key when the operation requires permission. Control handlers
  inherit `Authenticated` and `ControlReady` from `control.Register`.

## Style

- JSON field names use `snake_case`.
- Imports are formatted by `gci`, `goimports` and `golines`; keep lines at 80
  characters. Do not manually defeat the formatter.
- `wsl_v5` and `nlreturn` require readable whitespace between logical blocks.
- Prefer concise, context-specific package aliases in the central control
  registry when a nested package name would make an event registration wrap.
- Do not introduce a direct database client or a compatibility abstraction
  around `elum2b/services` without an explicit requirement.

## Validation

Run after Go changes:

```sh
GOLANGCI_LINT_CACHE=/private/tmp/platform-golangci-cache \
GOCACHE=/private/tmp/platform-go-build \
golangci-lint run ./...

GOCACHE=/private/tmp/platform-go-build go test ./...
git diff --check
```

Use `golangci-lint fmt ./...` and, when necessary,
`golangci-lint run --fix ./...` before the final lint run. The repository uses
`golangci-lint` v2 configuration in `.golangci.yml`.
