package server

import (
	"reflect"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	type args struct {
		route RouteDefinition
	}
	tests := []struct {
		name string
		args args
		want HttpHandlerWithoutReturn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleRequest(tt.args.route); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
