package gouter

import (
	"net/http"
	"strings"

	"github.com/kpango/glog"
)

type gouter struct {
	routes Routes
	router *http.ServeMux
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
	}
}

func (g *gouter) GetRouter() *http.ServeMux {
	return g.router
}

func (g *gouter) AddRoute(name, route, method string, handler http.HandlerFunc) *gouter {
	// TODO: store same route & different method pattern
	g.router.Handle(route, routing(method, glog.HTTPLogger(name, handler)))
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
		if strings.EqualFold(r.Method, method) {
			handler.ServeHTTP(w, r)
		}
	})
}
