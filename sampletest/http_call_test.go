package sampletest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMakeHTTPCall(t *testing.T) {

	tests := []struct {
		name    string
		server *httptest.Server
		expectedResponse *Response
		expectedErr error
	}{
		{
			name:             "success",
			server:           httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": 1, "name": "hay", "description": "desc"}`))
			})),
			expectedResponse: &Response{
				ID:          1,
				Name:        "hay",
				Description: "desc",
			},
			expectedErr:      nil,
		},
		{
			name:             "failed",
			server:           httptest.NewServer(http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`error`))
			})),
			expectedResponse: nil,
			expectedErr:      ErrBadStatusCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()
			resp, err := MakeHTTPCall(tt.server.URL)
			if !reflect.DeepEqual(resp, tt.expectedResponse) {
				t.Errorf("MakeHTTPCall() got = %v, want %v", resp, tt.expectedResponse)
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("MakeHTTPCall() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
		})
	}
}
