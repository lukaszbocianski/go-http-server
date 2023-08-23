package server

import (
	"net/http"
	"testing"
)

func TestInitialization(t *testing.T) {
	server := New()

	t.Run("Should create server object", func(t *testing.T) {
		if server == nil {
			t.Errorf("Server not created")
		}
	})

	t.Run("Should init routes", func(t *testing.T) {
		if server.GetRoutes() == nil {
			t.Errorf("routes not initialized")
		}
		if len(*server.GetRoutes()) != 0 {
			t.Errorf("routes not initialized")
		}
	})
}

func simpleRoutesMap(url, method string) *RoutesMap {
	return &RoutesMap{
		url: {Method: method, Function: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
			return nil, nil
		}},
	}
}

func TestRoutesOperations(t *testing.T) {
	t.Run("Should add simple route", func(t *testing.T) {
		server := New()
		server.RegisterRoutes(*simpleRoutesMap("/xyz", "get"))
		if len(*server.GetRoutes()) != 1 {
			t.Errorf("Expect 1 registered route, recieved %d", len(*server.GetRoutes()))
		}
	})

	t.Run("Should concatenate two route maps", func(t *testing.T) {
		server := New()
		server.RegisterRoutes(*simpleRoutesMap("/xyz", ""))
		server.RegisterRoutes(*simpleRoutesMap("/xyz/a", ""))
		if len(*server.GetRoutes()) != 2 {
			t.Errorf("Expect 2 registered routes, recieved %d", len(*server.GetRoutes()))
		}
	})
}
