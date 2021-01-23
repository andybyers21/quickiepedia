package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Form a new HTTP request passed to the handler.
	// arg1 = method, arg2 = route, arg3 = request body
	req, err := http.NewRequest("GET", "", nil)

	// In case of error formatting request, fail & stop the test
	if err != nil {
		t.Fatal(err)
	}

	// HTTP recorder to act as target for request
	recorder := httptest.NewRecorder()

	// Create HTTP handler
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to the handler
	hf.ServeHTTP(recorder, req)

	// Check status code
	if status := recorder.Code; status != httpStatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check response is expected
	expected := "Hello, World!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
