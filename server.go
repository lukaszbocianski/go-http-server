package server

import (
	"errors"
	"fmt"
	"golang.org/x/exp/maps"
	"net/http"
)

type HttpServer struct {
	routes RoutesMap
}

func New() *HttpServer {
	return &HttpServer{routes: map[string]RouteDefinition{}}
}

func (s *HttpServer) RegisterRoutes(routes map[string]RouteDefinition) {
	if s.routes == nil {
		s.routes = map[string]RouteDefinition{}
	}
	maps.Copy(s.routes, routes)
}

func (s *HttpServer) GetRoutes() *RoutesMap {
	return &s.routes
}

func (s *HttpServer) Start() error {
	for k, v := range s.routes {
		fmt.Printf("Registered route: [%s] %s\n", v.Method, k)
		http.HandleFunc(k, HandleRequest(v))
	}
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Errorf("Server closed\n")
	} else if err != nil {
		fmt.Errorf("Error starting server: %s\n", err)
	}
	if err != nil {
		return err
	}
	return nil
}
