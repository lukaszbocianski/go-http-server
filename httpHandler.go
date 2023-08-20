package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRequest(route RouteDefinition) HttpHandlerWithoutReturn {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != route.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		res, err := route.Function(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			fmt.Println(err.Error())
			return
		}
		if res != nil {
			json.NewEncoder(w).Encode(res)
		}
	}
}
