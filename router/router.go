package router

import "net/http"

type Route struct {
	Name string
	Path string
	Method string
	Handler func(w http.ResponseWriter, r *http.Request)
}

type IRouter interface {
	Append(route []Route) error
}

type router struct {
	routes []Route
}

func (router *router) Append(routes []Route) error {
	for _, route := range routes {
		router.routes = append(router.routes, route)
	}

	return nil
}

func New() IRouter {
	r := &router{}
	return r
}
