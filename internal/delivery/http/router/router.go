package router

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type contextKey string

const ParamKey contextKey = "pathParams"

type Route struct {
	Method      string
	Pattern     *regexp.Regexp
	ParamNames  []string
	HandlerFunc http.HandlerFunc
}

type prefixRoute struct {
	Method  string
	Prefix  string
	Handler http.Handler
}

type Router struct {
	routes       []Route
	prefixRoutes []prefixRoute
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	paramNames := []string{}
	regexPattern := regexp.MustCompile(`\{(\w+)\}`)
	replaced := regexPattern.ReplaceAllStringFunc(path, func(m string) string {
		name := m[1 : len(m)-1]
		paramNames = append(paramNames, name)
		return `([^/]+)`
	})

	finalRegex := regexp.MustCompile("^" + replaced + "(/.*)?$")

	r.routes = append(r.routes, Route{
		Method:      method,
		Pattern:     finalRegex,
		ParamNames:  paramNames,
		HandlerFunc: handler,
	})
}

func (r *Router) HandlePrefix(method, prefix string, handler http.Handler) {
	r.prefixRoutes = append(r.prefixRoutes, prefixRoute{
		Method:  method,
		Prefix:  prefix,
		Handler: handler,
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, pr := range r.prefixRoutes {
		if pr.Method == req.Method && strings.HasPrefix(req.URL.Path, pr.Prefix) {
			pr.Handler.ServeHTTP(w, req)
			return
		}
	}
	for _, route := range r.routes {
		if route.Method != req.Method {
			continue
		}
		matches := route.Pattern.FindStringSubmatch(req.URL.Path)
		if matches != nil {
			params := make(map[string]string)
			for i, name := range route.ParamNames {
				params[name] = matches[i+1]
			}
			ctx := context.WithValue(req.Context(), ParamKey, params)
			route.HandlerFunc(w, req.WithContext(ctx))
			return
		}
	}
	http.NotFound(w, req)
}

// GetParam retrieves param from context
func GetParam(r *http.Request, key string) string {
	params := r.Context().Value(ParamKey)
	if m, ok := params.(map[string]string); ok {
		return m[key]
	}
	return ""
}
