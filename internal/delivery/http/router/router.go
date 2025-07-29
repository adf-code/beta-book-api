package router

import "net/http"

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{
		routes: []Route{},
	}
}

func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: handler,
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	for _, route := range r.routes {
		if route.Method == method && matchPath(route.Pattern, path) {
			route.HandlerFunc.ServeHTTP(w, req)
			return
		}
	}

	http.NotFound(w, req)
}
