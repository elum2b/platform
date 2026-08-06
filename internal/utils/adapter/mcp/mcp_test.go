package mcp_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	serviceerrors "github.com/elum2b/services/errors"
	"github.com/gofiber/fiber/v3"
	mcpserver "github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/elum2b/platform/internal/api/methods"
	adapter "github.com/elum2b/platform/internal/utils/adapter"
	mcputils "github.com/elum2b/platform/internal/utils/adapter/mcp"
)

func registerMethods(router mcputils.Router) {
	methods.Register(adapter.Registry{MCP: router})
}

func TestRegisterCompilesAllTools(t *testing.T) {
	server := mcpserver.NewServer(
		&mcpserver.Implementation{Name: "test", Version: "test"},
		nil,
	)
	router := mcputils.NewRouter(server)

	registerMethods(router)
	router.Compile()
}

func TestInitAcceptsRequestsBeforeServicesStart(t *testing.T) {
	app := fiber.New()

	mcputils.Init(app, registerMethods)

	response, err := app.Test(httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/mcp",
		nil,
	))
	if err != nil {
		t.Fatalf("request MCP endpoint: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf(
			"status = %d, want %d",
			response.StatusCode,
			http.StatusBadRequest,
		)
	}
}

func TestToolsListIncludesAccessDiscovery(t *testing.T) {
	app := fiber.New()
	mcputils.Init(app, registerMethods)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/mcp",
		bytes.NewBufferString(
			`{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}`,
		),
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json, text/event-stream")
	request.Header.Set("MCP-Protocol-Version", "2025-06-18")

	response, err := app.Test(
		request,
		fiber.TestConfig{Timeout: 10 * time.Second},
	)
	if err != nil {
		t.Fatalf("request tools/list: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read tools/list response: %v", err)
	}

	for _, expected := range []string{
		"control.account.access.global.list",
		"control.account.access.workspace.list",
		"control.workspace.role.create",
		"control.workspace.invite.list",
		"control.workspace.audit.list",
		"control.global.role.permission.replace",
		"control.global.limit.resolve",
		"control.global.audit.list",
		"Moves the specified workspace to the archived state",
		`"anyOf"`,
		`"error"`,
	} {
		if !strings.Contains(string(body), expected) {
			t.Errorf("tools/list response does not contain %q", expected)
		}
	}
}

func TestToolErrorUsesStructuredContent(t *testing.T) {
	app := fiber.New()
	mcputils.Init(app, registerMethods)

	request := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/mcp",
		bytes.NewBufferString(
			`{"jsonrpc":"2.0","id":1,"method":"tools/call",`+
				`"params":{"name":"control.workspace.archive",`+
				`"arguments":{"workspace_id":`+
				`"00000000-0000-0000-0000-000000000000"}}}`,
		),
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json, text/event-stream")
	request.Header.Set("MCP-Protocol-Version", "2025-06-18")

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("request tools/call: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read tools/call response: %v", err)
	}

	for _, expected := range []string{
		`"content":[{"type":"text","text":"The MCP token has expired. ` +
			`Ask the user to issue a new MCP token before retrying."}]`,
		`"structuredContent":{"error":{"key":"UNAUTHORIZED",` +
			`"message":"unauthorized"}}`,
		`"isError":true`,
	} {
		if !strings.Contains(string(body), expected) {
			t.Errorf(
				"tools/call response does not contain %q: %s",
				expected,
				body,
			)
		}
	}
}

func TestForbiddenErrorInstructionNamesMethod(t *testing.T) {
	result := mcputils.ErrorResult(
		"control.workspace.archive",
		serviceerrors.ErrForbidden,
	)

	structured, ok := result.StructuredContent.(mcputils.ErrorResponse)
	if !ok {
		t.Fatalf(
			"structured content type = %T, want mcp.ErrorResponse",
			result.StructuredContent,
		)
	}

	if structured.Error.Key != serviceerrors.CodeForbidden {
		t.Errorf(
			"error key = %q, want %q",
			structured.Error.Key,
			serviceerrors.CodeForbidden,
		)
	}

	content, ok := result.Content[0].(*mcpserver.TextContent)
	if !ok {
		t.Fatalf("content type = %T, want *mcp.TextContent", result.Content[0])
	}

	want := "The account does not have permission to call " +
		"control.workspace.archive."
	if content.Text != want {
		t.Errorf("content text = %q, want %q", content.Text, want)
	}
}
