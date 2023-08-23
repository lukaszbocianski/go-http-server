package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func newRouteDef(method string, retVal interface{}, retError error) RouteDefinition {
	return RouteDefinition{
		Method: method,
		Function: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
			return retVal, retError
		},
	}
}

func marshalStruct(str interface{}) string {
	r, _ := json.Marshal(str)
	return string(r)
}

func TestHandleRequest(t *testing.T) {

	t.Run("Should not allow invalid method type and return 405", func(t *testing.T) {

		requestF := HandleRequest(newRouteDef(http.MethodPost, nil, nil))
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected code: %d , got: %d", http.StatusMethodNotAllowed, w.Code)
		}
	})

	t.Run("Should allow valid request type and return 200", func(t *testing.T) {
		requestF := HandleRequest(newRouteDef(http.MethodGet, nil, nil))
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected code: %d , got: %d", http.StatusOK, w.Code)
		}
	})

	t.Run("Should allow valid request type", func(t *testing.T) {
		requestF := HandleRequest(newRouteDef(http.MethodGet, nil, nil))
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected code: %d , got: %d", http.StatusOK, w.Code)
		}
		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected content type: 'application/json' , got: %s", w.Header().Get("Content-Type"))
		}
	})

	t.Run("Should return 400 and empty error for native golang errors", func(t *testing.T) {
		requestF := HandleRequest(newRouteDef(http.MethodGet, nil, http.ErrContentLength))
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected code: %d , got: %d", http.StatusBadRequest, w.Code)
		}
		if strings.TrimSpace(w.Body.String()) != "{}" {
			t.Errorf("Native golang exceptions should be translated to {}, not: '%s'", w.Body.String())
		}
	})

	t.Run("Should return 400 and error description for API errors", func(t *testing.T) {
		expectedError := ApiError{Code: "1", S: "testError"}
		requestF := HandleRequest(newRouteDef(http.MethodGet, nil, expectedError))
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		requestF(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected: %d , got: %d", http.StatusBadRequest, w.Code)
		}

		var returnedError ApiError
		json.Unmarshal(w.Body.Bytes(), &returnedError)

		if !reflect.DeepEqual(expectedError, returnedError) {
			t.Errorf("Expected '%s', got: '%s'", marshalStruct(expectedError), w.Body.String())
		}
	})
}
