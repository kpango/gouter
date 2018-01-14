package gouter

import (
	"context"
	"net/http"
	"strings"
	"sync"

	"github.com/kpango/gache"
	"github.com/kpango/glg"
)

type gouter struct {
	mu     sync.Mutex
	routes Routes
	router *http.ServeMux
	cache  *gache.Gache
	logFlg bool
}

type Path struct {
	Value  string
	Parent *Path
	Child  map[string]*Path
}

//Route structor
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func New() *gouter {
	return &gouter{
		routes: make([]Route, 1),
		router: http.NewServeMux(),
		cache:  gache.New(),
	}
}

func (g *gouter) GetRouter() *http.ServeMux {
	return g.router
}

func (g *gouter) EnableLogging() *gouter {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.logFlg = true
	glg.Get().SetMode(glg.NONE)
	return g
}

func (g *gouter) DisableLoggin() *gouter {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.logFlg = false
	glg.Get().SetMode(glg.NONE)
	return g
}

func (g *gouter) AddRoute(name, route, method string, handler http.HandlerFunc) *gouter {
	// TODO: store same route & different method pattern
	g.router.Handle(route, routing(method, glg.HTTPLogger(name, handler)))
	return g
}

func (g *gouter) SetRouter(Routes) *gouter {
	for _, route := range g.routes {
		g = g.AddRoute(route.Name, route.Pattern, route.Method, route.HandlerFunc)
	}
	return g
}

func routing(method string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.WithContext(context.WithValue(r.Context(), "key", "value"))
		if strings.EqualFold(r.Method, method) {
			handler.ServeHTTP(w, r)
		}
	})
}
