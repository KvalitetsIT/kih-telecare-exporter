package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthChecker(t *testing.T) {
	tests := []struct {
		method          string
		returnCode      int
		returns         int
		expectedFailure bool
	}{
		{method: http.MethodGet, returnCode: http.StatusOK, returns: http.StatusOK, expectedFailure: false},
		{method: http.MethodPost, returnCode: http.StatusOK, returns: http.StatusOK, expectedFailure: false},
		{method: http.MethodGet, returnCode: http.StatusOK, returns: http.StatusNotFound, expectedFailure: true},
		{method: http.MethodGet, returnCode: http.StatusOK, returns: http.StatusServiceUnavailable, expectedFailure: true},
		{method: http.MethodGet, returnCode: http.StatusOK, returns: http.StatusInternalServerError, expectedFailure: true},
		{method: http.MethodPost, returnCode: http.StatusOK, returns: http.StatusOK, expectedFailure: false},
	}

	for _, tt := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Setting return code: ", tt.returns)
			w.WriteHeader(tt.returns)
			w.Write([]byte(fmt.Sprintf("%d - returns", tt.returns))) //nolint
		}))
		defer ts.Close()

		client := http.Client{}

		if tt.expectedFailure {
			if err := PerformHealthCheck(client, tt.method, tt.returnCode, ts.URL); err == nil {
				t.Errorf("Expected failure - got no failure. return code - %v", err)
			}
		} else {
			if err := PerformHealthCheck(client, tt.method, tt.returnCode, ts.URL); err != nil {
				t.Errorf("Expected  no failure - got failure. return code - %v", err)
			}

		}
	}

}
