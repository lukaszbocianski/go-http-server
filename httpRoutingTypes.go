package server

import "net/http"

type RoutesMap map[string]RouteDefinition
type HttpHandler func(w http.ResponseWriter, r *http.Request) (interface{}, error)
type HttpHandlerWithoutReturn func(w http.ResponseWriter, r *http.Request)

type RouteDefinition struct {
	Method   string
	Function HttpHandler
}
