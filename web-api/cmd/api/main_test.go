package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home Handler",
			handler:        home,
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Health Handler",
			handler:        health,
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "About Handler",
			handler:        about,
			expectedStatus: http.StatusOK,
			expectedBody:   "Epi Peck",
		},
		{
			name:           "Time Handler",
			handler:        timeHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "", // just check it exists
		},
		{
			name:           "Greeting Handler",
			handler:        greeting,
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello user!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			tt.handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("got %v, expected %v", rr.Code, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("got %v, expected %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
