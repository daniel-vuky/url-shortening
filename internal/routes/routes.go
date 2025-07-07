package routes

import (
	"fmt"
	"net/http"

	"github.com/daniel-vuky/url-shortening/internal/handlers"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	mux *http.ServeMux
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// AddRoute adds a route to the router
func (r *Router) AddRoute(route Route) {
	r.mux.HandleFunc(fmt.Sprintf("%s /%s", route.Method, route.Path), route.Handler)
}

// AddGroup adds a group of routes to the router
func (r *Router) AddGroup(group string, routes []Route) {
	for _, route := range routes {
		r.mux.HandleFunc(fmt.Sprintf("%s /%s/%s", route.Method, group, route.Path), route.Handler)
	}
}

// GetMux returns the router's ServeMux
func (r *Router) GetMux() *http.ServeMux {
	return r.mux
}

func InitRoutes(handler *handlers.Handler) *Router {
	router := NewRouter()

	router.AddGroup(
		"/shorten",
		[]Route{
			{
				Path:    "/create",
				Method:  "POST",
				Handler: handler.CreateURL,
			},
		},
	)

	return router
}
