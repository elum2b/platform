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

- HTTP methods are under `internal/api/methods` and are registered through the
  universal adapter. HTTP routing, decoding and response envelopes belong in
  `internal/utils/adapter/http`.
- HTTP is reserved for authentication: provider login, two-factor completion,
  session check and TON challenge. Account logout and other authenticated work
  are socket methods.
- Keep session tokens only in cookies through `internal/utils/cookies`.
  Cookies must remain `Secure`, `HTTPOnly` and configured with `SameSite`.
- Provider-specific identity resolution belongs in `internal/utils/auth`, not
  in a controller.

## MCP API

- The MCP Streamable HTTP endpoint is `/mcp`. Routing, authentication, tool
  registration and protocol integration belong in
  `internal/utils/adapter/mcp`.
- Use the official `github.com/modelcontextprotocol/go-sdk/mcp` SDK. Do not
  implement MCP JSON-RPC framing manually.
- MCP authentication accepts `Authorization: Bearer mcp_...` and the
  `mcpToken` query parameter. Validate it through
  `services.Control.Internal.ValidateMCPToken` for every tool call.
  MCP tokens are not browser sessions and must never be read from cookies.
- Keep the transport stateless so revocation, account blocking and role changes
  take effect on the next MCP request. A tool must use its validated principal
  and the existing global/workspace access checks before calling a service.

## Socket API

- All non-authentication API methods are declared under `internal/api/methods`.
  ETP bootstrap, session authentication, system callbacks, decoding and
  responses belong in `internal/utils/adapter/socket`.
- Register transport-neutral methods through `internal/api/methods.Register`.
  A method declares its transports with `adapter.HTTP`, `adapter.WS` and
  `adapter.MCP`.
- Declare a method with this shape:

  ```go
  var Method = adapter.Method[Request, Response]{
      Key:         "service.section.method",
      Description: "Explain what the method does and when to call it.",
      Transports:  adapter.WS | adapter.MCP,
      Handler: func(ctx *adapter.Context, data Request) (Response, error) {
          // call services and return a transport-neutral response
      },
  }
  ```

- Methods are thin: validate transport input through the adapter, call
  `internal/services`, map a response and return service errors unchanged.
- One file contains one method. Keep request and response types in that same
  file; never create a shared `models.go` for a method section.
- Use nested packages for method domains, for example
  `account/identity/bind.go` or `workspace/role/permission/replace.go`.
- Add `adapter.GlobalAccess` or `adapter.WorkspaceAccess` using the exact
  service method key when the operation requires permission. WebSocket and MCP
  authentication are applied by their transport adapters.

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
