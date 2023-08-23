package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRequest(t *testing.T) {

	def := &RouteDefinition{
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
			return nil, nil
		},
	}
	requestF := HandleRequest(*def)

	t.Run("Should check allowed http method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected code: %d , got: %d", http.StatusMethodNotAllowed, w.Code)
		}
	})

	t.Run("Should allow valid request type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected code: %d , got: %d", http.StatusOK, w.Code)
		}
	})
}
