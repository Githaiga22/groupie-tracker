package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDateHandler(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		urlPath            string
		queryParams        string
		expectedStatusCode int
	}{
		{
			name:               "Valid Request",
			method:             http.MethodGet,
			urlPath:            "/dates",
			queryParams:        "?id=1",
			expectedStatusCode: http.StatusOK,
		},
		// {
		// 	name:               "Invalid Path",
		// 	method:             http.MethodGet,
		// 	urlPath:            "/invalid",
		// 	expectedStatusCode: http.StatusNotFound,
		// },
		// {
		// 	name:               "Invalid Method",
		// 	method:             http.MethodPost,
		// 	urlPath:            "/dates",
		// 	expectedStatusCode: http.StatusMethodNotAllowed,
		// },
		// {
		// 	name:               "Invalid ID - Out of Range",
		// 	method:             http.MethodGet,
		// 	urlPath:            "/dates",
		// 	queryParams:        "?id=100",
		// 	expectedStatusCode: http.StatusBadRequest,
		// },
		// {
		// 	name:               "Invalid ID - Non-numeric",
		// 	method:             http.MethodGet,
		// 	urlPath:            "/dates",
		// 	queryParams:        "?id=abc",
		// 	expectedStatusCode: http.StatusBadRequest,
		// },
		// {
		// 	name:               "Internal Server Error",
		// 	method:             http.MethodGet,
		// 	urlPath:            "/dates",
		// 	queryParams:        "?id=2", // Adjust based on your FetchDates simulation
		// 	expectedStatusCode: http.StatusInternalServerError,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the request
			req := httptest.NewRequest(tt.method, tt.urlPath+tt.queryParams, nil)
			w := httptest.NewRecorder()

			// Call the handler
			DateHandler(w, req)

			// Check the status code
			res := w.Result()
			if res.StatusCode != tt.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tt.expectedStatusCode, res.StatusCode)
			}
		})
	}
}
