package mcp

import (
	"context"
	"errors"
	"strings"

	mcpserver "github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	ErrRouteAlreadyExists = errors.New("mcp: route already exists")
	ErrRoutePatternEmpty  = errors.New("mcp: route pattern is empty")
	ErrMiddlewareNil      = errors.New("mcp: middleware is nil")
	ErrHandlerNil         = errors.New("mcp: handler is nil")
	ErrRouterCompiled     = errors.New("mcp: router is already compiled")
)

type Handler interface {
	register(*mcpserver.Server)
	Use(func(context.Context) (context.Context, error))
	UseInput(func(context.Context, any) error)
}

type Middleware func(Handler)

type Router interface {
	On(string, Handler) error
	Use(string, Middleware) error
	Group(...string) *Group
}

type middlewareRoute struct {
	pattern    string
	middleware Middleware
}

type endpointRoute struct {
	pattern string
	handler Handler
}

type router struct {
	server      *mcpserver.Server
	prefix      string
	parent      *router
	middlewares []middlewareRoute
	endpoints   []endpointRoute
	groups      []*router
	registered  map[string]struct{}
	compiled    bool
}

type App struct{ router *router }
type Group struct{ router *router }

func NewRouter(server *mcpserver.Server) *App {
	return &App{
		router: &router{server: server, registered: make(map[string]struct{})},
	}
}

func (app *App) Use(pattern string, middleware Middleware) error {
	return app.router.Use(pattern, middleware)
}

func (app *App) On(event string, handler Handler) error {
	return app.router.On(event, handler)
}

func (app *App) Group(prefix ...string) *Group {
	return &Group{router: app.router.Group(prefix...)}
}

func (app *App) Compile() { app.router.Compile() }

func (group *Group) Use(pattern string, middleware Middleware) error {
	return group.router.Use(pattern, middleware)
}

func (group *Group) On(event string, handler Handler) error {
	return group.router.On(event, handler)
}

func (group *Group) Group(prefix ...string) *Group {
	return &Group{router: group.router.Group(prefix...)}
}

func (r *router) Use(pattern string, middleware Middleware) error {
	if r.root().compiled {
		return ErrRouterCompiled
	}

	if middleware == nil {
		return ErrMiddlewareNil
	}

	r.middlewares = append(r.middlewares, middlewareRoute{
		pattern: r.scopedPattern(pattern), middleware: middleware,
	})

	return nil
}

func (r *router) On(event string, handler Handler) error {
	if r.root().compiled {
		return ErrRouterCompiled
	}

	if handler == nil {
		return ErrHandlerNil
	}

	pattern := joinPattern(r.fullPrefix(), event)
	if pattern == "" {
		return ErrRoutePatternEmpty
	}

	root := r.root()
	if _, ok := root.registered[pattern]; ok {
		return ErrRouteAlreadyExists
	}

	root.registered[pattern] = struct{}{}
	r.endpoints = append(
		r.endpoints,
		endpointRoute{pattern: pattern, handler: handler},
	)

	return nil
}

func (r *router) Group(prefix ...string) *router {
	if r.root().compiled {
		panic(ErrRouterCompiled)
	}

	if len(prefix) > 1 {
		panic("mcp: group accepts at most one prefix")
	}

	groupPrefix := ""
	if len(prefix) == 1 {
		groupPrefix = prefix[0]
	}

	group := &router{prefix: groupPrefix, parent: r}

	r.groups = append(r.groups, group)

	return group
}

func (r *router) Compile() {
	root := r.root()
	if root.compiled {
		return
	}

	root.compile(nil)

	root.compiled = true
}

func (r *router) compile(inherited []middlewareRoute) {
	middlewares := append(inherited, r.middlewares...)
	for _, endpoint := range r.endpoints {
		for _, middleware := range middlewares {
			if matchPattern(middleware.pattern, endpoint.pattern) {
				middleware.middleware(endpoint.handler)
			}
		}

		endpoint.handler.register(r.root().server)
	}

	for _, group := range r.groups {
		group.compile(middlewares)
	}
}

func (r *router) root() *router {
	if r.parent == nil {
		return r
	}

	return r.parent.root()
}

func (r *router) fullPrefix() string {
	if r.parent == nil {
		return r.prefix
	}

	return joinPattern(r.parent.fullPrefix(), r.prefix)
}

func (r *router) scopedPattern(pattern string) string {
	if pattern != "*" {
		return joinPattern(r.fullPrefix(), pattern)
	}

	prefix := r.fullPrefix()
	if prefix == "" {
		return pattern
	}

	return prefix + ".*"
}

func joinPattern(prefix string, pattern string) string {
	prefix = strings.Trim(prefix, ".")
	pattern = strings.Trim(pattern, ".")

	switch {
	case prefix == "":
		return pattern
	case pattern == "":
		return prefix
	default:
		return prefix + "." + pattern
	}
}

func matchPattern(pattern string, event string) bool {
	if pattern == "*" {
		return true
	}

	if strings.HasSuffix(pattern, ".*") {
		return strings.HasPrefix(event, strings.TrimSuffix(pattern, "*"))
	}

	return pattern == event
}
